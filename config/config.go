package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PrivateKey string
	Port       string
}

func GetConfig() Config {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error reading .env file")
	}

	var port string
	var isSet bool

	port, isSet = os.LookupEnv("PORT")

	if !isSet {
		port = "3000"
	}

	var config Config = Config{
		PrivateKey: os.Getenv("PRIVATE_KEY"),
		Port:       port,
	}

	return config
}
