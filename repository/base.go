package repository

import (
	"context"
	"github.com/sonymuhamad/todo-app/pkg"
	"gorm.io/gorm"
)

type BaseRepository struct {
	db *gorm.DB
}

func (r *BaseRepository) GetDB(ctx context.Context) *gorm.DB {
	if db, ok := ctx.Value(pkg.TrxDB).(*gorm.DB); ok {
		return db
	}

	return r.db.WithContext(ctx)
}
