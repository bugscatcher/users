package users

import (
	"context"
	"github.com/bugscatcher/users/services"
	"testing"

	"github.com/bugscatcher/users/application"
	"github.com/bugscatcher/users/config"
	"github.com/stretchr/testify/assert"
)

func TestGRPCHandler_GetUsers(t *testing.T) {
	conf, err := config.New()
	assert.NoError(t, err)
	app, err := application.New(&conf)
	assert.NoError(t, err)
	h := New(app)

	q := &services.QueryUsers{
		Search: "test",
	}
	result, err := h.GetUsers(context.Background(), q)
	assert.NoError(t, err)
	expectedResult := &services.ResultUsers{
		Users: nil,
	}
}
