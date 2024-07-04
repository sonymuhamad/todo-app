package repository

import (
	"context"
	
	"github.com/sonymuhamad/todo-app/model"
)

type UserRepository interface {
	GetByID(ctx context.Context, ID int64) (model.User, error)
}
