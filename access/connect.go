package access

import (
	"fmt"
	// "io/ioutil"
	"log"
	"net"

	// "path/filepath"
	"strings"

	"main/config"

	"golang.org/x/crypto/ssh"
)

func acceptAnyHostKey(_ string, _ net.Addr, _ ssh.PublicKey) error {
	return nil
}

func ConnectToServerWithPrivatePublicKeys(user string, host string, port string) {
	// private key parsing logic
	// exePath, err := os.Executable()

	// if err != nil {
	// 	panic(err)
	// }

	// privateKeyFilePath := []string{filepath.Dir(exePath), "keys/private-key.txt"}

	keys := config.GetConfig()

	privateKeyBytes := []byte(keys.PrivateKey)

	fmt.Println(keys.PrivateKey)

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
		HostKeyCallback: acceptAnyHostKey,
	}

	client, err := ssh.Dial("tcp", strings.Join([]string{host, port}, ":"), config)

	if err != nil {
		log.Fatal("Failed to dial: %s", err)
	}

	defer client.Close()

	session, err := client.NewSession()

	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}

	defer session.Close()

	output, err := session.CombinedOutput("pm2 status")

	if err != nil {
		log.Fatal("Failed to run command: %s", err)
	}

	fmt.Println(string(output))
}
