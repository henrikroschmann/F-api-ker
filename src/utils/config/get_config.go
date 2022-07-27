package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDb  string
	DATABASE string
}

func GetConfig() Config {
	var config Config
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config = Config{
		MongoDb:  os.Getenv("MONGODB_URL"),
		DATABASE: os.Getenv("DATABASE"),
	}
	return config
}
