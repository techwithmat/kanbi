package user

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

const InsertQuery = "INSERT INTO users (email,username,password) VALUES ($1, $2, $3) RETURNING id"
const GetByEmailQuery = "SELECT id, username, email, password FROM users WHERE email = $1"

type userRepo struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) userRepository {
	return &userRepo{
		db: db,
	}
}

func (repository *userRepo) Insert(ctx context.Context, user *RegisterRequest) (int, error) {
	var id int

	err := repository.db.QueryRow(ctx, InsertQuery, user.Email, user.Username, user.Password).Scan(&id)

	if err != nil {
		log.Println(err)

		return 0, err
	}

	return id, nil
}

func (repository *userRepo) GetByEmail(ctx context.Context, param string) (*User, error) {
	var (
		id                              int
		email, username, hashedPassword string
	)

	err := repository.db.QueryRow(ctx, GetByEmailQuery, param).Scan(&id, &username, &email, &hashedPassword)

	if err != nil {
		log.Println(err)

		return nil, err
	}

	return &User{
		ID:       id,
		Email:    email,
		Username: username,
		Password: hashedPassword,
	}, nil
}
