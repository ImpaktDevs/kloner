package main

import (
	"log"
	"os"

	"main/cmd"
)

func main() {
	app := cmd.Commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
