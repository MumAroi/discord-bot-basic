package config

import "os"

type Config struct {
	Prefix   string `json:"prefix"`
	BotToken string `json:"bot_token"`
	OwnerId  string `json:"owner_id"`
}

func LoadConfig() *Config {
	config := new(Config)
	config.BotToken = os.Getenv("DISCORD_TOKEN")
	config.Prefix = os.Getenv("PREFIX")
	return config
}
