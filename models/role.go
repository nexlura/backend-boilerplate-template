package models

type Role struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}
