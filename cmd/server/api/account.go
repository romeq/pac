package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/romeq/pac/pkg/db"
)

type createAccountModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *API) CreateAccount(ctx echo.Context) error {
	body, err := parseBodyToStruct[createAccountModel](ctx.Request().Body)
	if err != nil {
		return err
	}

	account, err := a.db.CreateAccount(ctx.Request().Context(), body.Username)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, account)
	return nil
}

type updateAccountRoleModel struct {
	User uuid.UUID `json:"username"`
	Role uuid.UUID `json:"role"`
}

func (a *API) AddAccountRole(ctx echo.Context) error {
	body, err := parseBodyToStruct[updateAccountRoleModel](ctx.Request().Body)
	if err != nil {
		return err
	}

	rec, err := a.db.AddRoleFor(ctx.Request().Context(), db.AddRoleForParams{
		AccountUuid: body.User,
		RoleUuid:    body.Role,
	})
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, rec)
	return nil
}
