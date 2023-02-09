package access

import (
	"bytes"
	"io/ioutil"
	"log"
	"net"
	"time"

	"strings"

	"github.com/ttacon/chalk"
	"golang.org/x/crypto/ssh"
)

func acceptAnyHostKey(_ string, _ net.Addr, _ ssh.PublicKey) error {
	return nil
}

const (
	PasswordAuth = "password-auth"
	KeyAuth      = "key-auth"
)

func ConnectToServerWithPrivatePublicKeys(user string, host string, port string, authType string, pkPath string, password string) {
	if pkPath == "" && authType == "key-auth" {
		log.Fatal(chalk.Red, "Please provide path to private key or change the authType to `password-auth`")
	}

	if password == "" && authType == "password-auth" {
		log.Fatal(chalk.Red, "Please provide a password or change the authType to `key-auth`")
	}

	var privateKey ssh.Signer

	if authType == "key-auth" {
		data, err := ioutil.ReadFile(pkPath)
		if err != nil {
			log.Fatal(chalk.Red, "Failed to load private key from file, please check file path")
		}

		privateKeyBytes := []byte(string(data))

		if privateKeyBytes == nil {
			log.Fatal(chalk.Red, "Failed to load private key")
		}

		privateKey, err = ssh.ParsePrivateKey(privateKeyBytes)

		if err != nil {
			log.Fatal(chalk.Red, "Failed to parse private key: %s", err)
		}
	}

	userNamePrivateKeyClientConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(privateKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
	}

	userNamePasswordClientConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
	}

	var config *ssh.ClientConfig

	switch authType {
	case PasswordAuth:
		config = userNamePasswordClientConfig
		break
	case KeyAuth:
		config = userNamePrivateKeyClientConfig
		break
	default:
		config = userNamePasswordClientConfig
		break
	}

	client, err := ssh.Dial("tcp", strings.Join([]string{host, port}, ":"), config)

	if err != nil {
		log.Fatal("Failed to dial: %s", err)
	}

	defer client.Close()

	go func() {
		for {
			// Open a session to run the `true` command
			session, err := client.NewSession()
			if err != nil {
				log.Fatalf("Failed to create session: %v", err)
			}
			defer session.Close()

			// Run the `true` command
			if err := session.Run("true"); err != nil {
				log.Fatalf("Failed to run: %v", err)
			}

			// Wait for a minute before running the command again
			time.Sleep(time.Minute)
		}
	}()

	session, err := client.NewSession()

	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}

	defer session.Close()

	var b bytes.Buffer

	session.Stdout = &b

	if err := session.Run("git version"); err != nil {
		log.Fatal(chalk.Red, "Failed to run command: %v", err)
	}

	log.Println(chalk.Green, string(b.String()))
}
