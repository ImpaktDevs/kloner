package mocker

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

// mockAuth is a simple authentication function that checks the provided username and password
func mockAuth(username, password string) bool {
	if username == "user" && password == "pass" {
		return true
	}
	return false
}

// startMockServer starts the mock SSH server and returns a function to shut it down
func StartMockServer() func() {
	ln, err := net.Listen("tcp", "localhost:2022")
	if err != nil {
		log.Fatal("Failed to start mock server:", err)
	}

	// Set up the SSH server config
	config := &ssh.ServerConfig{
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			if mockAuth(c.User(), string(pass)) {
				return nil, nil
			}
			return nil, fmt.Errorf("incorrect username or password")
		},
	}

	// Start the mock server
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Print("Failed to accept connection:", err)
				continue
			}

			_, chans, reqs, err := ssh.NewServerConn(conn, config)
			if err != nil {
				log.Print("Failed to create SSH server connection:", err)
				continue
			}

			go ssh.DiscardRequests(reqs)
			go func(chans <-chan ssh.NewChannel) {
				for newChannel := range chans {
					channel, requests, err := newChannel.Accept()
					if err != nil {
						log.Print("Failed to accept channel:", err)
						continue
					}

					go func(in <-chan *ssh.Request) {
						for req := range in {
							req.Reply(false, nil)
						}
					}(requests)

					channel.Close()
				}
			}(chans)
		}
	}()

	// Return a function to shut down the mock server
	return func() {
		ln.Close()
	}
}
