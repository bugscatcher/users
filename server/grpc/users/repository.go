package users

import (
	"github.com/bugscatcher/users/models"
	"github.com/jackc/pgx"
	"golang.org/x/xerrors"
)

func update(pool *pgx.ConnPool) {
	pool.Exec(`
		UPDATE users
		SET first_name = $1 AND last_name = $2
		WHERE id = ""`)
}

func findUser(pool *pgx.ConnPool, id []string) ([]*models.User, error) {
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
			id = ANY($1)`, id)
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
