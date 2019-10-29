package main

import (
	"database/sql"
	"fmt"

	"github.com/bugscatcher/users/config"
	_ "github.com/bugscatcher/users/migration/common/migrations"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/rs/zerolog/log"
	"golang.org/x/xerrors"
)

func main() {
	conf, err := config.New()
	if err != nil {
		xErr := xerrors.Errorf("while reading config: %w", err)
		log.Fatal().Err(xErr).Msg("Read config")
	}
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", conf.PostgreSQL.Host, conf.PostgreSQL.User, conf.PostgreSQL.Database)
	migrationsConn, err := sql.Open("postgres", connStr)
	if err != nil {
		xErr := xerrors.Errorf("while opening db connection: %w", err)
		log.Fatal().Err(xErr).Msg("Open DB connection")
	}
	if err := goose.SetDialect("postgres"); err != nil {
		xErr := xerrors.Errorf("while setting dialect: %w", err)
		log.Fatal().Err(xErr).Msg("Set dialect")
	}
	if err := goose.Up(migrationsConn, "."); err != nil {
		xErr := xerrors.Errorf("while migrating up: %w", err)
		log.Fatal().Err(xErr).Msg("Migrate up")
	}
	_ = migrationsConn.Close()
}
