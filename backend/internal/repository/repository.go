package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	GetUserByEmail(string) (*models.User, error)
	InsertUser(models.User) (int, error)
}
