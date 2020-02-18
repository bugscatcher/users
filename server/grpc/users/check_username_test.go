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

func TestHandler_CheckUsername(t *testing.T) {
	h := newTestHandler(0)
	user := testutil.GetRandomUser(h.userID)
	err := addUsers(h.db, user)
	assert.NoError(t, err)
	testCases := []struct {
		name     string
		username string
		expResp  *services.CheckUsernameResult
	}{
		{
			"check if valid username is already taken (upper case)",
			strings.ToUpper(user.Username),
			&services.CheckUsernameResult{Result: &services.CheckUsernameResult_IsAvailable{IsAvailable: false}},
		},
		{
			"check if valid username is already taken (lower case)",
			strings.ToLower(user.Username),
			&services.CheckUsernameResult{Result: &services.CheckUsernameResult_IsAvailable{IsAvailable: false}},
		},
		{
			"check if valid username is already taken (both case)",
			strings.ToLower(user.Username)[:1] + strings.ToUpper(user.Username)[1:],
			&services.CheckUsernameResult{Result: &services.CheckUsernameResult_IsAvailable{IsAvailable: false}},
		},
		{
			"check if valid username isn't taken",
			testutil.GetRandomUsername(),
			&services.CheckUsernameResult{Result: &services.CheckUsernameResult_IsAvailable{IsAvailable: true}},
		},
		{
			"check if username starts with space",
			" " + faker.Username(),
			&services.CheckUsernameResult{Result: &services.CheckUsernameResult_Error{Error: ErrorMessageInvalidUsername}},
		},
		{
			"check if username starts with number",
			"9" + faker.Username(),
			&services.CheckUsernameResult{Result: &services.CheckUsernameResult_Error{Error: ErrorMessageUsernameStartWithNumber}},
		},
		{
			fmt.Sprintf("check if username length less than %d", MinUsernameLength),
			faker.Username()[:MinUsernameLength-1],
			&services.CheckUsernameResult{Result: &services.CheckUsernameResult_Error{Error: ErrorMessageUsernameLength}},
		},
		{
			fmt.Sprintf("check if username length equal %d", MinUsernameLength),
			"a" + strconv.Itoa(time.Now().UTC().Nanosecond())[:MinUsernameLength-1],
			&services.CheckUsernameResult{Result: &services.CheckUsernameResult_IsAvailable{IsAvailable: true}},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			req := &services.CheckUsernameRequest{Username: tc.username}
			res, err := h.service.CheckUsername(h.ctx, req)
			assert.NoError(t, err)
			assert.EqualValues(t, tc.expResp, res)
			actUser, err := findUsers(h.db, h.userID.String())
			assert.ElementsMatch(t, []*models.User{user}, actUser)
		})
	}
}
