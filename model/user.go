package model

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Username    string    `json:"username"`
	PhoneNumber string    `json:"phone_number"`
}
