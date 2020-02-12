package headers

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

func AddUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDHeader, userID)
}

func GetUserID(ctx context.Context) (uuid.UUID, error) {
	s, ok := ctx.Value(userIDHeader).(string)
	if !ok {
		return uuid.Nil, errors.New("value is not type of string")
	}
	userID, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, errors.New("value can't be parsed to UUID")
	}
	return userID, nil
}
