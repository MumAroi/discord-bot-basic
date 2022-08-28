package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Prefix   string
	BotToken string
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func LoadConfig() *Config {
	config := new(Config)
	config.BotToken = os.Getenv("DISCORD_TOKEN")
	config.Prefix = os.Getenv("PREFIX")
	return config
}
