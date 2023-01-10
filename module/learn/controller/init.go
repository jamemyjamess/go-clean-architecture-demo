package controller

import (
	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/usecase"
	"github.com/labstack/echo/v4"
)

type SessionController interface {
	SetSession(c echo.Context) error
	GetSession(c echo.Context) error
}

type sessionController struct {
	UserUsecase usecase.UserUsecase
}

func NewSessionController(e *echo.Group, userUsecase usecase.UserUsecase) {
	sessionControllers := &sessionController{
		UserUsecase: userUsecase,
	}
	e.Group("/learn-session")
	{
		e.GET("/:id", sessionControllers.SetSession)
		e.PUT("/:id", sessionControllers.GetSession)
	}
}

func NewSessionHandler(userUsecase usecase.UserUsecase) SessionController {
	return &sessionController{userUsecase}
}
