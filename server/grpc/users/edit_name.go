package users

import (
	"context"
	"fmt"

	"github.com/bugscatcher/users/headers"
	"github.com/bugscatcher/users/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) EditName(ctx context.Context, req *services.RequestEditName) (*services.Response, error) {
	userID, err := headers.GetUserID(ctx)
	if err != nil {
		return nil, status.Error(codes.Canceled, fmt.Sprintf("problem with userID: %v", err))
	}
	err = update(h.db, req.FirstName, req.LastName, userID.String())
	if err != nil {
		return nil, err
	}
	return &services.Response{Status: services.Status_OK}, nil
}
