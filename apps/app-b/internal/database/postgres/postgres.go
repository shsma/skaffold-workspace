package postgres

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

const latestMigrationVersion = 25

func init() {
	const connectionStr = "postgres://localhost:5432/database?sslmode=enable"
	runMigrate(connectionStr)
}

func runMigrate(connectionStr string) {
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to open SQL connection: %v", connectionStr)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)
	err = m.Migrate(latestMigrationVersion)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to run migration: %v", latestMigrationVersion)
	}
}
