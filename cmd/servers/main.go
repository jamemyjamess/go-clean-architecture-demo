package main

import (
	"github.com/jamemyjamess/go-clean-architecture-demo/app"
	"github.com/jamemyjamess/go-clean-architecture-demo/configs"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	configs.Init(e)
	// migrateRouters.Init(e)
	app := app.NewApp()
	app.Run()
	// e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
