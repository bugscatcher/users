package users

import (
	"context"

	headers "github.com/bugscatcher/go-deps"
	"github.com/bugscatcher/users/services"
)

func (h *Handler) EditName(ctx context.Context, req *services.RequestEditName) (*services.Response, error) {
	userID, err := headers.GetUserID(ctx)
	if err != nil {
		return nil, getGRPCErrorForGetUserIDHeader(err)
	}
	err = update(h.db, req.FirstName, req.LastName, userID.String())
	if err != nil {
		return nil, err
	}
	return &services.Response{Status: services.Status_OK}, nil
}
