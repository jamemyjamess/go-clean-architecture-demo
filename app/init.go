package app

import (
	"os"

	"github.com/jamemyjamess/go-clean-architecture-demo/configs"
	"github.com/labstack/echo/v4"
)

type App struct {
	Echo *echo.Echo
	// Cfg *configs.Configs
	// Db *gorm.DB
}

func NewApp() *App {
	return &App{
		Echo: echo.New(),
		// Db:   database.PostgresSql,
	}
}

func (app *App) Run() {
	// e := echo.New()
	configs.Init(app.Echo)
	// migrateRouters.Init(e)
	app.Echo.Logger.Fatal(app.Echo.Start(":" + os.Getenv("PORT")))
}
