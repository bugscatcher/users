package users

import (
	"context"

	"github.com/bugscatcher/users/models"
	"github.com/bugscatcher/users/services"
	"github.com/google/uuid"
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
	return &services.ResultUsers{Users: models.ToUsers(users...)}, nil
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
