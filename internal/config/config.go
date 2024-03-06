package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Grpc struct {
		Port    int    `yaml:"port"`
		Timeout string `yaml:"timeout"`
	}
}

func ReadConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}