package models

import (
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
	Status    string    `json:"status"` // verified, pending, disabled
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProfileDto struct {
	Id        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ProfileFromDomainList(pf []*Profile) []ProfileDto {
	output := make([]ProfileDto, len(pf))
	for i, p := range pf {
		output[i] = ProfileFromDomain(p)
	}
	return output
}

func ProfileFromDomain(p *Profile) ProfileDto {
	return ProfileDto{
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
