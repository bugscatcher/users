package users

import (
	"fmt"

	"github.com/bugscatcher/users/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"golang.org/x/xerrors"
)

func update(pool *pgx.ConnPool, firstName, lastName, id string) error {
	cmdTag, err := pool.Exec(`
		UPDATE users
		SET first_name = $1, last_name = $2
		WHERE id = $3
	`, firstName, lastName, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return fmt.Errorf("no row foud to update")
	}
	return nil
}

func findUsers(pool *pgx.ConnPool, ids ...string) ([]*models.User, error) {
	result := make([]*models.User, 0)
	rows, err := pool.Query(`
		SELECT
			id,
			first_name,
			last_name,
			username
		FROM
			users
		WHERE
			id = ANY($1)`, ids)
	if err != nil {
		return nil, xerrors.Errorf("while selecting from users: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		row := &models.User{}
		err := rows.Scan(&row.ID, &row.FirstName, &row.LastName, &row.Username)
		if err != nil {
			return nil, xerrors.Errorf("while scanning row: %w", err)
		}
		result = append(result, row)
	}
	return result, nil
}

func addUsers(pool *pgx.ConnPool, user ...*models.User) error {
	_, err := pool.CopyFrom(
		pgx.Identifier{"users"},
		[]string{"id", "first_name", "last_name", "username"},
		pgx.CopyFromRows(models.GetValues(user...)),
	)
	return err
}

func isUsernameAvailable(pool *pgx.ConnPool, username string) (bool, error) {
	var isAvailable bool
	err := pool.QueryRow(`SELECT NOT EXISTS(SELECT 1 FROM users WHERE username = LOWER($1))`, username).Scan(&isAvailable)
	return isAvailable, err
}

func updateUsername(pool *pgx.ConnPool, username string, id uuid.UUID) error {
	_, err := pool.Exec(`
		UPDATE users 
		SET username = $1 
		WHERE id = $2
	`, username, id)
	if err != nil {
		return err
	}
	return nil
}
