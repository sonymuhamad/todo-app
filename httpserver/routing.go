package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/sonymuhamad/todo-app/pkg"
)

func InitRouter(cfg pkg.EnvConfig, h *Handler) *echo.Echo {
	r := echo.New()

	r.GET("/", h.Index)
	r.GET("/user", h.GetUserByID)

	return r
}
