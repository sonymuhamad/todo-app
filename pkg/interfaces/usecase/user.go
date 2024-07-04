package usecase

import (
	"context"
	
	"github.com/sonymuhamad/todo-app/model"
)

type UserUsecase interface {
	GetByID(ctx context.Context, ID int64) (model.User, error)
}
