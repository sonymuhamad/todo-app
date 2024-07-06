package usecase

import (
	"context"
	"github.com/sonymuhamad/todo-app/httpserver"

	"github.com/go-playground/validator/v10"
	"github.com/sonymuhamad/todo-app/model"
	"github.com/sonymuhamad/todo-app/pkg/interfaces/repository"
	"github.com/sonymuhamad/todo-app/pkg/interfaces/usecase"
)

type User struct {
	userRepository repository.UserRepository
	validator      *validator.Validate
}

func NewUser(userRepository repository.UserRepository,
	validator *validator.Validate) *User {
	return &User{
		userRepository: userRepository,
		validator:      validator,
	}
}

func (u *User) GetByID(ctx context.Context, ID int64) (model.User, error) {

	return u.userRepository.GetByID(ctx, ID)
}

func (u *User) Create(ctx context.Context, param usecase.CreateUserParam) error {
	err := u.validator.Struct(param)
	if err != nil {
		return err
	}

	return httpserver.UnprocessableError{
		Field:   "test error",
		Message: "nahlo error",
	}

	//return nil
}
