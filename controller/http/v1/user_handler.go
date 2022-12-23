package handler

import "github.com/labstack/echo/v4"

func (handler *Handler) mapUserRouteHandler(e *echo.Group) {
	e.Group("/company")
	{
		e.GET("/:id", handler.user.FindInfo)
		e.PUT("/:id", handler.user.CreateOrUpdate)
	}
}
