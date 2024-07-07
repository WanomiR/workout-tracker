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

	query := `SELECT id, email, password, name, patronymic, surname,
       CASE WHEN weight IS NULL THEN 0 ELSE weight END,
       CASE WHEN height IS NULL THEN 0 ELSE height END,
       CASE WHEN dob IS NULL THEN '1970-01-01'::DATE ELSE dob END,
       registered_at
FROM public.users WHERE email = $1`

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return 0, errors.New("error hashing password")
	}

	if user.DateOfBrith == "" {
		user.DateOfBrith = "1970-01-01" // replace empty value with default date to avoid errors
	}

	query := `INSERT INTO public.users (email, password, name, patronymic, surname, weight, height, dob, registered_at) 
VALUES ($1, $2, $3, $4, $5, NULLIF($6, 0), NULLIF($7, 0), NULLIF($8, '1970-01-01'::DATE), now()) RETURNING id`

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
