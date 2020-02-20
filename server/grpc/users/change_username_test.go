package users

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/bugscatcher/users/models"
	"github.com/bugscatcher/users/services"
	"github.com/bugscatcher/users/testutil"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestHandler_ChangeUsername(t *testing.T) {
	h := newTestHandler(0)
	user := testutil.GetRandomUser(h.userID)
	err := addUsers(h.db, user)
	assert.NoError(t, err)
	testCases := []struct {
		name     string
		username string
		expResp  *services.Response
	}{
		{
			"can't change if valid username is already taken (upper case)",
			strings.ToUpper(user.Username),
			&services.Response{Status: services.Status_ALREADY_EXISTS},
		},
		{
			"can't change if valid username is already taken (lower case)",
			strings.ToLower(user.Username),
			&services.Response{Status: services.Status_ALREADY_EXISTS},
		},
		{
			"can't change if valid username is already taken (both case)",
			strings.ToLower(user.Username)[:1] + strings.ToUpper(user.Username)[1:],
			&services.Response{Status: services.Status_ALREADY_EXISTS},
		},
		{
			"can change if valid username isn't taken",
			testutil.GetRandomUsername(),
			&services.Response{Status: services.Status_OK},
		},
		{
			"can't change if username starts with space",
			" " + faker.Username(),
			&services.Response{Status: services.Status_INVALID_ARGUMENT},
		},
		{
			"can't change if username starts with number",
			"9" + faker.Username(),
			&services.Response{Status: services.Status_INVALID_ARGUMENT},
		},
		{
			fmt.Sprintf("can't change if username length less than %d", MinUsernameLength),
			faker.Username()[:MinUsernameLength-1],
			&services.Response{Status: services.Status_INVALID_ARGUMENT},
		},
		{
			fmt.Sprintf("can change if username length equal %d", MinUsernameLength),
			"a" + strconv.Itoa(time.Now().UTC().Nanosecond())[:MinUsernameLength-1],
			&services.Response{Status: services.Status_OK},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := h.service.ChangeUsername(h.ctx, &services.UsernameRequest{Username: tc.username})
			assert.NoError(t, err)
			assert.EqualValues(t, tc.expResp, res)
			if tc.expResp.Status == services.Status_OK {
				user.Username = tc.username
			}
			actUser, err := findUsers(h.db, h.userID.String())
			assert.EqualValues(t, []*models.User{user}, actUser)
		})
	}
}
