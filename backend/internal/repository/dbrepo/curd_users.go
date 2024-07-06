package dbrepo

import (
	"backend/internal/models"
	"context"
)

func (db *PostgresDbRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM public.users WHERE email = $1`

	var user models.User
	err := db.Conn.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Patronymic,
		&user.Surname,
		&user.Weight,
		&user.Height,
		&user.DateOfBrith,
		&user.RegisteredAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
