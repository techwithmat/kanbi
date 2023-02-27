package user

import "context"

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

type RegisterRequest struct {
	Email                string `json:"email" validate:"required,email,min=6,max=32"`
	Username             string `json:"username" validate:"required,min=3,max=32"`
	Password             string `json:"password" validate:"required,eqfield=PasswordConfirmation,min=8,max=28"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,min=8,max=28"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,min=6,max=32"`
	Password string `json:"password" validate:"required,min=8,max=28"`
}

type DeleteRequest struct {
	Password string `json:"password" validate:"required,min=8,max=28"`
}

type userRepository interface {
	Insert(ctx context.Context, user *RegisterRequest) (int, error)
	GetByEmail(ctx context.Context, param string) (*User, error)
}
