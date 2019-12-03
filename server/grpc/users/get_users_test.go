package users

import (
	"context"
	"testing"

	"github.com/bugscatcher/users/model"
	"github.com/bugscatcher/users/services"
	"github.com/bugscatcher/users/testutil"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetUsers_OneUser(t *testing.T) {
	h := newTestHandler(0)
	user := testutil.GetRandomUser()
	err := addUsers(h.db, user)
	assert.NoError(t, err)
	req := &services.RequestGetUsers{Id: []string{user.ID.String()}}
	result, err := h.service.GetUsers(context.Background(), req)
	assert.NoError(t, err)
	expectedResult := &services.ResultUsers{Users: model.ToUsers(user)}
	assert.EqualValues(t, expectedResult, result)
}

func TestHandler_GetUsers_MultipleUsers(t *testing.T) {
	h := newTestHandler(0)
	users := testutil.GetRandomUsers(5)
	err := addUsers(h.db, users...)
	assert.NoError(t, err)
	req := &services.RequestGetUsers{
		Id: []string{
			users[0].ID.String(),
			users[1].ID.String(),
			users[2].ID.String(),
		},
	}
	result, err := h.service.GetUsers(context.Background(), req)
	assert.NoError(t, err)
	expectedResult := &services.ResultUsers{Users: model.ToUsers(users[0:3]...)}
	assert.EqualValues(t, expectedResult, result)
}

func TestHandler_GetUsers_NegativeCases(t *testing.T) {
	tests := []struct {
		name string
		req  *services.RequestGetUsers
	}{
		{
			"empty request",
			&services.RequestGetUsers{},
		},
		{
			"empty ids",
			&services.RequestGetUsers{
				Id: []string{},
			},
		},
		{
			"id isn't uuid",
			&services.RequestGetUsers{
				Id: []string{"string"},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			h := newTestHandler(0)
			result, err := h.service.GetUsers(context.Background(), tc.req)
			assert.Error(t, err)
			assert.Nil(t, result)
		})
	}
}

func TestHandler_GetUsers_MultipleUsers_ExistentAndNonExistent(t *testing.T) {
	h := newTestHandler(0)
	users := testutil.GetRandomUsers(5)
	err := addUsers(h.db, users...)
	assert.NoError(t, err)
	nonexistentID := uuid.New()
	req := &services.RequestGetUsers{
		Id: []string{
			users[0].ID.String(),
			users[4].ID.String(),
			users[2].ID.String(),
			nonexistentID.String(),
		},
	}
	result, err := h.service.GetUsers(context.Background(), req)
	assert.NoError(t, err)
	expectedUsers := model.ToUsers([]*model.User{users[0], users[4], users[2]}...)
	assert.ElementsMatch(t, expectedUsers, result.Users)
}

func addUsers(pool *pgx.ConnPool, user ...*model.User) error {
	_, err := pool.CopyFrom(
		pgx.Identifier{"users"},
		[]string{"id", "first_name", "last_name", "username"},
		pgx.CopyFromRows(model.GetValues(user...)),
	)
	return err
}
