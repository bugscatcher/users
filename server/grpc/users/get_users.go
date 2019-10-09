package users

import (
	"context"

	"github.com/bugscatcher/users/services"
)

func (G *GRPCHandler) GetUsers(context.Context, *services.QueryUsers) (*services.ResultUsers, error) {
	panic("implement me")
}
