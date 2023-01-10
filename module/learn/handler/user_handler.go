package handler

import (
	"github.com/jamemyjamess/go-clean-architecture-demo/module/learn/controller"
	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/repository/postgres"
	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/usecase"
	"github.com/jamemyjamess/go-clean-architecture-demo/pkg/database"
	"github.com/labstack/echo/v4"
)

func NewLearnHandler(e *echo.Group) {
	usersRepository := postgres.NewUserRepository(database.PostgresSql)
	usersUsecase := usecase.NewUserUsecase(usersRepository, nil)
	controller.NewSessionController(e, usersUsecase)
}
