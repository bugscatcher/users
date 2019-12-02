package testutil

import (
	"github.com/bugscatcher/users/model"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

func GetRandomUser() *model.User {
	return &model.User{
		ID:        uuid.New(),
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Username:  faker.Username(),
	}
}

func GetRandomUsers(n int) []*model.User {
	result := make([]*model.User, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, GetRandomUser())
	}
	return result
}
