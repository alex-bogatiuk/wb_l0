package repo

import (
	_ "context"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/jackc/pgx/v5"
)

type Config struct {
	DBProvider string `yaml:"db_provider"`
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
		return &cfg, err
	}

	return &cfg, nil
}
