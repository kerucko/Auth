package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Grpc            GrpcConfig     `yaml:"grpc"`
	Database        DatabaseConfig `yaml:"database"`
	ToketExpiration time.Duration  `yaml:"toket_expiration"`
}

type GrpcConfig struct {
	Port    int    `yaml:"port"`
	Timeout string `yaml:"timeout"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Timeout  string `yaml:"timeout"`
}

func MustReadConfig() Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("config path is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config not found: " + configPath)
	}
	var config Config
	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		panic("config not read: " + err.Error())
	}
	return config
}
