package access

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"

	"strings"

	"golang.org/x/crypto/ssh"
)

func acceptAnyHostKey(_ string, _ net.Addr, _ ssh.PublicKey) error {
	return nil
}

func ConnectToServerWithPrivatePublicKeys(user string, host string, port string, pkPath string, pkType string) {
	data, err := ioutil.ReadFile(pkPath)

	if err != nil {
		log.Fatal("Failed to load private key from file, please check file path")
	}

	privateKeyBytes := []byte(string(data))

	if privateKeyBytes == nil {
		log.Fatal("Failed to load private key")
	}

	privateKey, err := ssh.ParsePrivateKey(privateKeyBytes)

	if err != nil {
		log.Fatal("Failed to parse private key: %s", err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(privateKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
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
		log.Fatal("Failed to run command: %v", err)
	}

	fmt.Println(string(b.String()))
}
