package config

import (
	"log"
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
		log.Fatal("Error reading .env file: %s", err)
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
