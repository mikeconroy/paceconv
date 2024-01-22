package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mikeconroy/paceconv/view"
)

type IndexHandler struct{}

func (h *IndexHandler) HandleIndex(c echo.Context) error {
	page := view.ShowIndex()
	return page.Render(c.Request().Context(), c.Response())
}
