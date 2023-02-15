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

	"main/parser"
)

func acceptAnyHostKey(_ string, _ net.Addr, _ ssh.PublicKey) error {
	return nil
}

const (
	PasswordAuth = "password-auth"
	KeyAuth      = "key-auth"
)

func ConnectToServerWithPrivatePublicKeys(user string, host string, port string, authType string, pkPath string, password string) *ssh.Session {
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

	var b bytes.Buffer

	session.Stdout = &b

	defer session.Close()

	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}

	return session
}

func runSteps(step parser.Process, client *ssh.Client, info []string, errors []string, successes []string) {
	log.Println(chalk.Magenta, strings.Join([]string{":::::", step.Description, ":::::"}, ""), chalk.Reset)
	info = append(info, strings.Join([]string{":::::", step.Description, ":::::"}, ""))
	for _, item := range step.Commands {
		session, err := client.NewSession()

		defer session.Close()

		if err != nil {
			log.Fatalf("Failed to create session: %s", err)
		}
		var b bytes.Buffer

		session.Stdout = &b
		log.Println(chalk.Green, strings.Join([]string{">", item}, " "), chalk.Reset)
		if err := session.Run(item); err != nil {
			log.Println(chalk.Red, strings.Join([]string{err.Error()}, ": "), chalk.Reset)
			errors = append(errors, strings.Join([]string{err.Error()}, ": "))
		} else {
			log.Println(strings.Join([]string{b.String()}, ": "))
			successes = append(successes, strings.Join([]string{b.String()}, ": "))
			b.Reset()
		}

	}
}

func HandleWorkflowInServer(path string, target string) {
	workflow := parser.ParseWorkflowFile(path)

	user, host, port, authType, pkPath, password := workflow.ServerInfo.User, workflow.ServerInfo.Host, workflow.ServerInfo.Port, workflow.ServerInfo.AuthType, workflow.ServerInfo.PrivateKeyPath, workflow.ServerInfo.Password

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

	errors := []string{}
	successes := []string{}
	info := []string{}

	targetValue, exists := workflow.Targets[target]
	if !exists || target == "" {
		log.Fatal(chalk.Red, "No target picked", chalk.Reset)
	}

	log.Println(chalk.Magenta, strings.Join([]string{"::::: Running Pre process", ":::::"}, " "), chalk.Reset)
	if len(workflow.PreProcess) > 0 {
		for _, step := range workflow.PreProcess {
			runSteps(step, client, info, errors, successes)
		}
	} else {
		log.Println(chalk.Blue, "No pre process commands", chalk.Reset)
	}

	log.Println(chalk.Magenta, strings.Join([]string{"::::: Running target", target, ":::::"}, " "), chalk.Reset)
	for _, step := range targetValue {
		runSteps(step, client, info, errors, successes)
	}

	log.Println(chalk.Magenta, strings.Join([]string{"::::: Running Post process", ":::::"}, " "), chalk.Reset)
	if len(workflow.PostProcess) > 0 {
		for _, step := range workflow.PostProcess {
			runSteps(step, client, info, errors, successes)
		}
	} else {
		log.Println(chalk.Blue, "No post process commands", chalk.Reset)
	}
}
