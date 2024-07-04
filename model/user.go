package model

import (
	"time"

	"github.com/sonymuhamad/todo-app/pkg/enum"
)

type User struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	Role      enum.UserRole `json:"role"`
	IsActive  bool          `json:"is_active"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type UserSession struct {
	ID           int64      `json:"id"`
	AccessToken  string     `json:"access_token"`
	UserID       int64      `json:"user_id"`
	LastSeenTime time.Time  `json:"last_seen_time"`
	ExpiredAt    *time.Time `json:"expired_at"`
	LoginTime    time.Time  `json:"login_time"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
