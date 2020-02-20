package users

import (
	"context"

	"github.com/bugscatcher/users/headers"
	"github.com/bugscatcher/users/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getGRPCErrorForGetUserIDHeader(err error) error {
	return status.Error(codes.Canceled, err.Error())
}

func (h *Handler) ChangeUsername(ctx context.Context, in *services.UsernameRequest) (*services.Response, error) {
	userID, err := headers.GetUserID(ctx)
	if err != nil {
		return nil, getGRPCErrorForGetUserIDHeader(err)
	}
	result, err := checkUsername(h.db, in.Username)
	if err != nil {
		return nil, err
	}
	switch res := result.Result.(type) {
	case *services.CheckUsernameResult_IsAvailable:
		if res.IsAvailable {
			if err := updateUsername(h.db, in.Username, userID); err != nil {
				return nil, err
			}
		} else {
			return &services.Response{Status: services.Status_ALREADY_EXISTS}, nil
		}

	case *services.CheckUsernameResult_Error:
		return &services.Response{Status: services.Status_INVALID_ARGUMENT}, nil
	}
	return &services.Response{Status: services.Status_OK}, nil
}
