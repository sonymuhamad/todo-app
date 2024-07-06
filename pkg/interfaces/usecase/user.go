package usecase

import (
	"context"
	"github.com/sonymuhamad/todo-app/pkg/enum"

	"github.com/sonymuhamad/todo-app/model"
)

type UserUsecase interface {
	GetByID(ctx context.Context, ID int64) (model.User, error)
	Create(ctx context.Context, param CreateUserParam) error
}

type CreateUserParam struct {
	Name            string        `json:"name" validate:"required,min=1"`
	Email           string        `json:"email" validate:"required,min=1"`
	Role            enum.UserRole `json:"role"`
	Password        string        `json:"password" validate:"required,min=1"`
	ConfirmPassword string        `json:"confirm_password" validate:"required,min=1"`
}
