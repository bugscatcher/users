package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upInit, downInit)
}

var initDB = `
CREATE TABLE users (
	id uuid PRIMARY KEY,
	first_name text,
	last_name text,
	username text
);

CREATE TABLE users_phone (
	user_id uuid PRIMARY KEY,
	phone_number text,
	FOREIGN KEY (user_id) REFERENCES users (id)
);
`

func upInit(tx *sql.Tx) error {
	_, err := tx.Exec(initDB)
	if err != nil {
		return err
	}
	return nil
}

func downInit(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
