package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type createRoleModel struct {
	Name string `json:"name"`
}

func (a *API) CreateRole(ctx echo.Context) error {
	body, err := parseBodyToStruct[createRoleModel](ctx.Request().Body)
	if err != nil {
		return err
	}

	account, err := a.db.CreateRole(ctx.Request().Context(), body.Name)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, account)
	return nil
}
