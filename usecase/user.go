package usecase

import (
	"context"
	"github.com/sonymuhamad/todo-app/model"
	"github.com/sonymuhamad/todo-app/pkg/interfaces/repository"
)

type User struct {
	userRepository repository.UserRepository
}

func NewUser(userRepository repository.UserRepository) *User {
	return &User{
		userRepository: userRepository,
	}
}

func (u *User) GetByID(ctx context.Context, ID int64) (model.User, error) {
	return u.userRepository.GetByID(ctx, ID)
}
