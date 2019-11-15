package users

import (
	"context"
	"github.com/bugscatcher/users/model"
	"github.com/bugscatcher/users/services"
	"github.com/bxcodec/faker/v3"
	"github.com/jackc/pgx"
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
	user := &model.User{}
	err = faker.FakeData(&user)
	assert.NoError(t, err)
	err = addUser(h.db, user)
	assert.NoError(t, err)
	q := &services.QueryUsers{
		Search: user.FirstName,
	}
	result, err := h.GetUsers(context.Background(), q)
	assert.NoError(t, err)
	expectedResult := &services.ResultUsers{
		Users: []*services.User{toGRPCUser(user)},
	}
	assert.EqualValues(t, expectedResult, result)
}

func addUser(pool *pgx.ConnPool, user *model.User) error {
	_, err := pool.Exec(`
		INSERT INTO users (id, first_name, last_name, username)
		VALUES($1, $2, $3, $4)`, user.ID, user.FirstName, user.LastName, user.Username)
	return err
}

func toGRPCUser(user *model.User) *services.User {
	return &services.User{
		Id:          user.ID.String(),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Username:    user.Username,
	}
}
