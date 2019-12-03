package users

import (
	"context"

	"github.com/bugscatcher/users/model"
	"github.com/bugscatcher/users/services"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) GetUsers(ctx context.Context, req *services.RequestGetUsers) (*services.ResultUsers, error) {
	if err := checkRequestGetUsers(req); err != nil {
		return nil, err
	}
	users, err := findUser(h.db, req.Id)
	if err != nil {
		return nil, err
	}
	return &services.ResultUsers{Users: model.ToUsers(users...)}, nil
}

func findUser(pool *pgx.ConnPool, id []string) ([]*model.User, error) {
	result := make([]*model.User, 0)
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
		row := &model.User{}
		err := rows.Scan(&row.ID, &row.FirstName, &row.LastName, &row.Username)
		if err != nil {
			return nil, xerrors.Errorf("while scanning row: %w", err)
		}
		result = append(result, row)
	}
	return result, nil
}

func checkRequestGetUsers(req *services.RequestGetUsers) error {
	if len(req.Id) == 0 {
		return status.Error(codes.InvalidArgument, "ID is empty")
	}
	for _, id := range req.Id {
		if !isUUID(id) {
			return status.Error(codes.InvalidArgument, "ID is isn't uuid")
		}
	}
	return nil
}

func isUUID(s string) bool {
	_, err := uuid.Parse(s)
	if err != nil {
		return false
	}
	return true
}
