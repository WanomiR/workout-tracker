package dbrepo

import (
	"backend/internal/models"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
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

func (db *PostgresDbRepo) InsertUser(user models.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `INSERT INTO public.users (email, password, name, patronymic, surname, weight, height, dob, registered_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now()) RETURNING id`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, errors.New("error hashing password")
	}

	var userId int
	err = db.Conn.QueryRowContext(ctx, query,
		user.Email,
		string(hashedPassword),
		user.Name,
		user.Patronymic,
		user.Surname,
		user.Weight,
		user.Height,
		user.DateOfBrith,
	).Scan(&userId)

	if err != nil {
		return 0, err
	}

	return userId, nil
}
