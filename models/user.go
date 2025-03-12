package models

import (
	"mime/multipart"
	"time"
)

type Profile struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	RoleId    string    `json:"role_id"`
	Status    string    `json:"status"` // active, pending, disabled
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProfileFrom struct {
	ID        string                `form:"id"`
	FirstName string                `form:"first_name"`
	LastName  string                `form:"last_name"`
	Email     string                `form:"email"`
	Password  string                `form:"password"`
	Phone     string                `form:"phone"`
	RoleId    string                `form:"role_id"`
	Status    string                `form:"status"` // active, pending, disabled
	Avatar    *multipart.FileHeader `form:"avatar"`
	CreatedAt time.Time             `form:"created_at"`
	UpdatedAt time.Time             `form:"updated_at"`
}

type ProfileDTO struct {
	Id        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ProfileFromDomainList(pf []Profile) []ProfileDTO {
	output := make([]ProfileDTO, len(pf))
	for i, p := range pf {
		output[i] = ProfileFromDomain(p)
	}
	return output
}

func ProfileFromDomain(p Profile) ProfileDTO {
	return ProfileDTO{
		Id:        p.ID,
		FirstName: p.FirstName,
		Email:     p.Email,
		LastName:  p.LastName,
		Avatar:    p.Avatar,
		Phone:     p.Phone,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
