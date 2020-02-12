package models

import (
	"github.com/bugscatcher/users/services"
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Username    string    `json:"username"`
	PhoneNumber string    `json:"phone_number"`
}

func (m *User) toUser() *services.User {
	return &services.User{
		Id:          m.ID.String(),
		FirstName:   m.FirstName,
		LastName:    m.LastName,
		Username:    m.Username,
		PhoneNumber: m.PhoneNumber,
	}
}

func ToUsers(users ...*User) []*services.User {
	result := make([]*services.User, 0, len(users))
	for _, user := range users {
		result = append(result, user.toUser())
	}
	return result
}

func (m *User) getValues() []interface{} {
	return []interface{}{m.ID, m.FirstName, m.LastName, m.Username}
}

func GetValues(users ...*User) [][]interface{} {
	result := make([][]interface{}, 0, len(users))
	for _, user := range users {
		result = append(result, user.getValues())
	}
	return result
}
