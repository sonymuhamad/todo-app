package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/sonymuhamad/todo-app/pkg"
)

func InitRouter(cfg pkg.EnvConfig, h *Handler) *echo.Echo {
	r := echo.New()
	r.HTTPErrorHandler = CustomHTTPErrorHandler

	r.GET("/", h.Index)
	r.GET("/user", h.GetUserByID)
	r.POST("/user", h.CreateUser)

	return r
}
