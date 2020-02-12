package testutil

import (
	"github.com/bugscatcher/users/models"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

func GetRandomUser() *models.User {
	return &models.User{
		ID:        uuid.New(),
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Username:  faker.Username(),
	}
}

func GetRandomUsers(n int) []*models.User {
	result := make([]*models.User, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, GetRandomUser())
	}
	return result
}
