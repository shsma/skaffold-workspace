package config

import (
	"github.com/gosidekick/goconfig"
	"github.com/rs/zerolog/log"
)

type CoreConfig struct {
}

type ServerConfig struct {
	Port string `yaml:"http_server_port" cfg:"http_server_port"`
}

type DatabaseConfig struct {
	DbDriver   string `yaml:"db_driver" cfg:"db_driver"`
	DbHost     string `yaml:"db_host" cfg:"db_host"`
	DbPort     string `yaml:"db_port" cfg:"db_port"`
	DbSslmode  string `yaml:"db_sslmode" cfg:"db_sslmode"`
	DbName     string `yaml:"db_name" cfg:"db_name"`
	DbUsername string `yaml:"db_username" cfg:"db_username"`
	DbPassword string `yaml:"db_password" cfg:"db_password"`
}

func LoadServerConfig() ServerConfig {
	cfg := ServerConfig{}
	err := goconfig.Parse(&cfg)
	if err != nil {
		log.Error().Err(err).Msg("Failed to load ServerConfig")
	}
	return cfg
}

func (cfg *ServerConfig) LogServerConfig() {
	log.Info().Msgf("%+v", cfg)
}

func LoadDatabaseConfig() DatabaseConfig {
	cfg := DatabaseConfig{}
	err := goconfig.Parse(&cfg)
	if err != nil {
		log.Error().Err(err).Msg("Failed to load DatabaseConfig")
	}
	return cfg
}

func (cfg *DatabaseConfig) LogDatabaserConfig() {
	log.Info().Msgf("%+v", cfg)
}
