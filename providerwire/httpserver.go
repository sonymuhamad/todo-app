//go:build wireinject
// +build wireinject

package providerwire

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/sonymuhamad/todo-app/httpserver"
	"github.com/sonymuhamad/todo-app/pkg"
)

func InitializeHttpServer() *echo.Echo {
	wire.Build(
		pkg.LoadEnvConfig,
		provideGormSQLDatabase,
		provideRepository,
		provideUsecase,
		httpserver.NewHandlerWithWire,
		provideHttpServer,
		//httpserver.InitRouter,
	)

	return &echo.Echo{}
}

func provideHttpServer(cfg pkg.EnvConfig, h *httpserver.Handler) *echo.Echo {
	r := httpserver.InitRouter(cfg, h)

	return r
}
