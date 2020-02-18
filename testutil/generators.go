package testutil

import (
	"strconv"
	"strings"
	"time"

	"github.com/bugscatcher/users/models"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

func GetRandomUsername() string {
	return strings.ToLower(faker.Username()) + strconv.Itoa(time.Now().UTC().Nanosecond())
}

func GetRandomUser(id uuid.UUID) *models.User {
	return &models.User{
		ID:        id,
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Username:  GetRandomUsername(),
	}
}

func GetRandomUsers(n int) []*models.User {
	result := make([]*models.User, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, GetRandomUser(uuid.New()))
	}
	return result
}
