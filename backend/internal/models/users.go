package models

import "time"

type User struct {
	ID           int       `json:"-"`
	Email        string    `json:"email" example:"admin@example.com" binding:"required"`
	Password     string    `json:"password" example:"password" binding:"required"` // internally is stored as hash
	Name         string    `json:"name" example:"John" binding:"required"`
	Patronymic   string    `json:"patronymic" example:"Ivanovich"`
	Surname      string    `json:"surname" example:"Doe"`
	Weight       float64   `json:"weight" example:"104.5"`
	Height       float64   `json:"height" example:"196.5"`
	DateOfBrith  string    `json:"dob" example:"1995-01-21"`
	RegisteredAt time.Time `json:"-"`
}

type UserCredentials struct {
	Email    string `json:"email" example:"admin@example.com" binding:"required"`
	Password string `json:"password" example:"password" binding:"required"` // internally is stored as hash
}
