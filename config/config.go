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
		log.Fatal("Error reading .env file")
	}

	var config Config = Config{
		PrivateKey: os.Getenv("PRIVATE_KEY"),
		Port:       os.Getenv("PORT"),
	}

	return config
}
