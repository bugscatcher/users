package users

import (
	"context"
	"testing"

	headers "github.com/bugscatcher/go-deps"
	"github.com/bugscatcher/users/models"
	"github.com/bugscatcher/users/services"
	"github.com/bugscatcher/users/testutil"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHandler_EditName(t *testing.T) {
	testCases := []struct {
		name string
		req  *services.RequestEditName
	}{
		{
			"user can change his first name and last name",
			&services.RequestEditName{FirstName: faker.FirstName(), LastName: faker.LastName()},
		},
		{
			"user can change his first name (to random) and last name (to empty)",
			&services.RequestEditName{FirstName: faker.FirstName()},
		},
		{
			"user can change his first name (to empty) and last name (to random)",
			&services.RequestEditName{LastName: faker.LastName()},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			h := newTestHandler(0)
			user := testutil.GetRandomUser(h.userID)
			err := addUsers(h.db, user)
			assert.NoError(t, err)
			resp, err := h.service.EditName(h.ctx, tc.req)
			assert.NoError(t, err)
			assert.EqualValues(t, &services.Response{Status: services.Status_OK}, resp)
			actUser, err := findUsers(h.db, h.userID.String())
			assert.NoError(t, err)
			user.FirstName = tc.req.FirstName
			user.LastName = tc.req.LastName
			assert.ElementsMatch(t, []*models.User{user}, actUser)
		})
	}
}

func TestHandler_EditName_Negative(t *testing.T) {
	h := newTestHandler(0)
	testCases := []struct {
		name string
		ctx  context.Context
	}{
		{
			"can't change name if context is empty",
			context.Background(),
		},
		{
			"can't change name if user isn't exist",
			headers.AddUserID(context.Background(), uuid.New().String()),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			resp, err := h.service.EditName(tc.ctx, &services.RequestEditName{
				FirstName: faker.FirstName(), LastName: faker.LastName(),
			})
			assert.Nil(t, resp)
			assert.Error(t, err)
		})
	}
}
