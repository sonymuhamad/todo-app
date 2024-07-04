package repository

import (
	"context"
	"errors"
	"github.com/sonymuhamad/todo-app/model"
	"gorm.io/gorm"
)

type User struct {
	BaseRepository
}

func NewUser(db *gorm.DB) *User {
	return &User{
		BaseRepository{
			db: db,
		},
	}
}

func (u *User) GetByID(ctx context.Context, ID int64) (model.User, error) {
	db := u.GetDB(ctx)

	var user model.User

	err := db.Where("id = ?", ID).
		Preload("Patients").
		Preload("Patients.PatientCredit").
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("error bosku")
		}

		return user, err
	}

	return user, nil
}
