package httpserver

import (
	"context"
	"github.com/labstack/echo/v4"
	usecaseInterface "github.com/sonymuhamad/todo-app/pkg/interfaces/usecase"
	"net/http"
)

type Handler struct {
	userUsecase usecaseInterface.UserUsecase
}

func NewHandlerWithWire(userUsecase usecaseInterface.UserUsecase) *Handler {
	return &Handler{
		userUsecase: userUsecase,
	}
}

func (h *Handler) Index(ctx echo.Context) error {
	data := "Hello from /index"
	return ctx.String(http.StatusOK, data)
}

func (h *Handler) GetUserByID(ctx echo.Context) error {
	ctxNew := context.Background()
	user, err := h.userUsecase.GetByID(ctxNew, int64(10))
	if err != nil {
		return err
	}

	return ctx.String(http.StatusOK, user.Name)
}
