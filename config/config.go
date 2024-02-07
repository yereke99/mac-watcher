package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SSHUser     string
	SSHPassword string

	DBConfig    *DatabaseConfig
	Password    string
	SlackBotUrl string
	ChannelName string
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	Database string
}

func NewConfig(fileName string) (*Config, error) {
	conf := new(Config)

	loadEnvFile(fileName)

	databaseConfig := &DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	}

	conf.DBConfig = databaseConfig
	conf.SSHUser = os.Getenv("SSH_USER")
	conf.SSHPassword = os.Getenv("SSH_PASSWORD")
	conf.SlackBotUrl = os.Getenv("SLACK_BOT_URL")
	conf.ChannelName = os.Getenv("CHANNEL_NAME")

	return conf, nil
}

func loadEnvFile(fileName string) {

	if err := godotenv.Load(fileName); err != nil {
		log.Fatal(err)
	}
}
