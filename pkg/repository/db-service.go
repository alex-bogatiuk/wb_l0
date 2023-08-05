package repo

import (
	"github.com/gookit/slog"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DBProvider string `yaml:"dbProvider"`
	Host       string `yaml:"host"`
	Basename   string `yaml:"basename"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Port       string `yaml:"port"`
}

func InitConfig() (*Config, error) {
	var cfg Config

	// Read configuration file. Settings from the ...\configs\cfg.yml file
	err := cleanenv.ReadConfig("configs/cfg.yml", &cfg)
	if err != nil {
		slog.Error("cfg file reading error")

		return &cfg, err
	}

	return &cfg, nil
}
