//go:build wireinject
// +build wireinject

package providerwire

import (
	"github.com/google/wire"
	repositoryInterface "github.com/sonymuhamad/todo-app/pkg/interfaces/repository"
	"github.com/sonymuhamad/todo-app/repository"
)

var repositorySet = wire.NewSet(repository.NewUser)

var provideRepository = wire.NewSet(
	repositorySet,
	wire.Bind(new(repositoryInterface.UserRepository), new(*repository.User)),
)
