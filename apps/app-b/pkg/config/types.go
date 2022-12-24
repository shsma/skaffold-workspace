package config

import (
	"github.com/gosidekick/goconfig"
	"github.com/rs/zerolog/log"
)

type CoreConfig struct {
}

type ServerConfig struct {
	Port     string         `yaml:"http_server_port" cfg:"http_server_port"`
	Database DatabaseConfig `yaml:"database" cfg:"database"`
}

type DatabaseConfig struct {
	DbDriver   string `yaml:"db_driver" cfg:"db_driver"`
	DbHost     string `yaml:"db_host" cfg:"db_host"`
	DbSslmode  string `yaml:"db_sslmode" cfg:"db_sslmode"`
	DbName     string `yaml:"db_name" cfg:"db_name"`
	DbUsername string `yaml:"db_username" cfg:"db_username"`
	DbPassword string `yaml:"db_password" cfg:"db_password"`
}

func LoadServerConfig() *ServerConfig {
	cfg := ServerConfig{}
	err := goconfig.Parse(&cfg)
	if err != nil {
		log.Error().Err(err).Msg("Failed to load ServerConfig")
	}
	dbCfg := DatabaseConfig{}
	err = goconfig.Parse(&dbCfg)
	if err != nil {
		log.Error().Err(err).Msg("Failed to load DatabaseConfig")
	}

	cfg.Database = dbCfg
	return &cfg
}

func (cfg *ServerConfig) LogServerConfig() {
	log.Info().Msgf("%+v", cfg)
}
