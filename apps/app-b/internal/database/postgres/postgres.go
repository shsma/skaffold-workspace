package postgres

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/shsma/app-b/pkg/config"
	"path"
	"runtime"
)

const latestMigrationVersion = 25

func init() {
	cfg := config.LoadDatabaseConfig()
	err := runMigrate(cfg)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to init DB: ", err)
		return
	}
}

func runMigrate(cfg config.DatabaseConfig) error {
	connectionStr := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DbDriver,
		cfg.DbUsername,
		cfg.DbPassword,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
		cfg.DbSslmode,
	)

	db, err := sql.Open(cfg.DbDriver, connectionStr)
	err = db.Ping()
	if err != nil {
		log.Error().Err(err).Msgf("Failed to ping the DB: ", connectionStr)
		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("%s/migrations/", workdir()), cfg.DbName, driver)
	if err != nil {
		log.Error().Err(err).Msgf("%s", connectionStr)
		log.Error().Err(err).Msgf("%s", fmt.Sprintf("%s/migrations/", workdir()))
		log.Error().Err(err).Msg("Failed to create a new Migrate instance")
		return err
	}
	err = m.Migrate(latestMigrationVersion)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to run migration: %v", latestMigrationVersion)
		return err
	}
	log.Info().Msg("DB migrations ran successfully")
	return nil
}

func workdir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename)
}
