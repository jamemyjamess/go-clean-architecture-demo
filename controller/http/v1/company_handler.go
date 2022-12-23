package handler

import "github.com/labstack/echo/v4"

func (handler *Handler) mapCompanyRouteHandler(e *echo.Group) {
	e.Group("/company")
	{
		e.GET("/:id", handler.company.Find)
	}
}
