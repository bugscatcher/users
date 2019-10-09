package postgresql

import (
	"time"

	"github.com/jackc/pgx"
)

type Config struct {
	Host             string        `mapstructure:"host"`
	Port             int           `mapstructure:"port"`
	Database         string        `mapstructure:"database"`
	User             string        `mapstructure:"user"`
	Password         string        `mapstructure:"password"`
	MaxConnections   int           `mapstructure:"max_connections"`
	MaxConnectionAge time.Duration `mapstructure:"max_connection_age"`
	Secured          bool          `mapstructure:"secured"`
}

func NewConnPool(config *Config) (*pgx.ConnPool, error) {
	db, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     config.Host,
			Port:     uint16(config.Port),
			Database: config.Database,
			User:     config.User,
			Password: config.Password,
		},
		MaxConnections: config.MaxConnections,
	})
	return db, err
}
