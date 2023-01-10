package handler

import "github.com/labstack/echo/v4"

func (handler *Handler) mapLearnRouteHandler(e *echo.Group) {
	e.Group("/learn")
	{
		e.GET("-session/setsession/:id", handler.learn.SetSession)
		e.PUT("-session/getsession/:id", handler.learn.GetSession)
	}
}
