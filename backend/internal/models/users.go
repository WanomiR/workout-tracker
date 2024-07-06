package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email" example:"admin@example.com" binding:"required"`
	Password     string    `json:"password" example:"password" binding:"required"` // internally is stored as hash
	Name         string    `json:"name" example:"Admin" binding:"required"`
	Patronymic   string    `json:"patronymic" example:"patronymic"`
	Surname      string    `json:"surname" example:"surname"`
	Weight       float64   `json:"weight" example:"104.5"`
	Height       float64   `json:"height" example:"196.5"`
	DateOfBrith  time.Time `json:"dob" example:"1995-01-21"`
	RegisteredAt time.Time `json:"-"`
}

type UserCredentials struct {
	Email    string `json:"email" example:"admin@example.com" binding:"required"`
	Password string `json:"password" example:"password" binding:"required"` // internally is stored as hash
}
