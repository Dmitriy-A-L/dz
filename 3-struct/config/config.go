package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func Load() *Config {
	err := godotenv.Load("key.env")
	if err != nil {
		log.Println("Файл key.env не найден")
	}

	return &Config{
		Key: os.Getenv("KEY1"),
	}
}
