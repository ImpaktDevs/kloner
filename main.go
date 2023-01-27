package main

import (
	"main/access"
	"main/server"
)

func main() {
	access.ConnectToServerWithPrivatePublicKeys("tora", "18.185.110.231", "22")

	server.StartServer()
}
