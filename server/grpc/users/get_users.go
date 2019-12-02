package users

import (
	"context"

	"github.com/bugscatcher/users/model"
	"github.com/bugscatcher/users/services"
	"github.com/jackc/pgx"
	"golang.org/x/xerrors"
)

func (h *Handler) GetUsers(ctx context.Context, q *services.QueryUsers) (*services.ResultUsers, error) {
	users, err := findUser(h.db, q)
	if err != nil {
		return &services.ResultUsers{}, err
	}
	result := make([]*services.User, 0)
	for _, user := range users {
		u := &services.User{
			Id:        user.ID.String(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
		}
		result = append(result, u)
	}
	return &services.ResultUsers{
		Users: result,
	}, nil
}

func findUser(pool *pgx.ConnPool, q *services.QueryUsers) ([]*model.User, error) {
	result := make([]*model.User, 0)
	//TODO search by first_name, last_name, username
	rows, err := pool.Query(`
		SELECT
			id,
			first_name,
			last_name,
			username
		FROM
			users
		WHERE
			first_name LIKE '%' || $1 || '%'`, q.Search)
	if err != nil {
		return nil, xerrors.Errorf("while selecting from users: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		row := &model.User{}
		err := rows.Scan(&row.ID, &row.FirstName, &row.LastName, &row.Username)
		if err != nil {
			return nil, xerrors.Errorf("while scanning row: %w", err)
		}
		result = append(result, row)
	}
	return result, nil
}
