package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *API) NotFound(ctx echo.Context) error {
	ctx.JSON(http.StatusNotFound, errorResponse{
		Error: "Not found",
	})
	return nil
}
