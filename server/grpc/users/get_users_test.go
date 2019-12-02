package users

import (
	"context"
	"github.com/bugscatcher/users/model"
	"github.com/bugscatcher/users/services"
	"github.com/bugscatcher/users/testutil"
	"github.com/jackc/pgx"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHandler_GetUsers_OneUser(t *testing.T) {
	h := newTestHandler(0)
	user := testutil.GetRandomUser()
	err := addUser(h.db, user)
	assert.NoError(t, err)
	q := &services.QueryUsers{
		Search: user.FirstName,
	}
	result, err := h.service.GetUsers(context.Background(), q)
	assert.NoError(t, err)
	expectedResult := &services.ResultUsers{Users: model.ToUsers(user)}
	assert.EqualValues(t, expectedResult, result)
}

func TestHandler_GetUsers_MultipleUsers(t *testing.T) {
	h := newTestHandler(0)
	users := testutil.GetRandomUsers(5)
	suffix := time.Now().String()
	for _, user := range users {

	}
	err := addUsers(h.db, users...)
	assert.NoError(t, err)
	q := &services.QueryUsers{
		Search: users[0].FirstName,
	}
	result, err := h.service.GetUsers(context.Background(), q)
	assert.NoError(t, err)
	expectedResult := &services.ResultUsers{Users: model.ToUsers(users[0])}
	assert.EqualValues(t, expectedResult, result)
}

func addUser(pool *pgx.ConnPool, user *model.User) error {
	_, err := pool.Exec(`
		INSERT INTO users (id, first_name, last_name, username)
		VALUES($1, $2, $3, $4)`, user.ID, user.FirstName, user.LastName, user.Username)
	return err
}

func addUsers(pool *pgx.ConnPool, user ...*model.User) error {
	_, err := pool.CopyFrom(
		pgx.Identifier{"users"},
		[]string{"id", "first_name", "last_name", "username"},
		pgx.CopyFromRows(model.GetValues(user...)),
	)
	return err
}
