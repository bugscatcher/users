package users

import (
	"context"
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/bugscatcher/users/services"
	"github.com/jackc/pgx"
)

const (
	MinUsernameLength                   int    = 5
	ErrorMessageInvalidUsername         string = "Sorry, this username is invalid."
	ErrorMessageUsernameStartWithNumber string = "Sorry, an username can't start with a number."
	ErrorMessageRegexp                  string = "Something went wrong :( We couldn't check username right now, please try again later."
)

var (
	ErrorMessageUsernameLength = fmt.Sprintf("An username must have at least %d characters.", MinUsernameLength)
)

func (h *Handler) CheckUsername(ctx context.Context, in *services.UsernameRequest) (*services.CheckUsernameResult, error) {
	return checkUsername(h.db, in.Username)
}

func checkUsername(pool *pgx.ConnPool, username string) (*services.CheckUsernameResult, error) {
	result := &services.CheckUsernameResult{}
	if err := validateUsername(username); err != nil {
		result.Result = err
		return result, nil
	}
	isAvailable, err := isUsernameAvailable(pool, username)
	if err != nil {
		return nil, err
	}
	result.Result = &services.CheckUsernameResult_IsAvailable{IsAvailable: isAvailable}
	return result, nil
}

func validateUsername(username string) *services.CheckUsernameResult_Error {
	result := &services.CheckUsernameResult_Error{}
	pattern := "^[0-9].*"
	matched, err := regexp.Match(pattern, []byte(username))
	if err != nil {
		result.Error = ErrorMessageRegexp
		return result
	}
	if matched {
		result.Error = ErrorMessageUsernameStartWithNumber
		return result
	}
	if utf8.RuneCountInString(username) < MinUsernameLength {
		result.Error = ErrorMessageUsernameLength
		return result
	}
	pattern = "^[A-Za-z][A-Za-z0-9]*"
	matched, err = regexp.Match(pattern, []byte(username))
	if err != nil {
		result.Error = ErrorMessageRegexp
		return result
	}
	if !matched {
		result.Error = ErrorMessageInvalidUsername
		return result
	}
	return nil
}
