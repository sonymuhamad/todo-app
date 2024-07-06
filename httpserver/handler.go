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

	return ctx.JSON(http.StatusOK, user)
}

func (h *Handler) CreateUser(c echo.Context) error {
	var user usecaseInterface.CreateUserParam
	ctx := context.Background()

	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	err := h.userUsecase.Create(ctx, user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, "Success create user")
}
