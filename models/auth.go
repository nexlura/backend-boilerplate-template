package models

import "time"

type Login struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	RoleId    string    `json:"role_id"`
	Password  string    `json:"password"`
	Status    string    `json:"status"` // verified, pending, disabled
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
