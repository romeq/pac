package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/romeq/pac/pkg/db"
)

type newResourceModel struct {
	Content         string      `json:"content"`
	AuthorizedRoles []uuid.UUID `json:"authorizedRoles"`
}

func (a *API) CreateResource(ctx echo.Context) error {
	body, err := parseBodyToStruct[newResourceModel](ctx.Request().Body)
	if err != nil {
		return err
	}

	resource, err := a.db.CreateResource(ctx.Request().Context(), body.Content)
	if err != nil {
		return err
	}

	for _, role := range body.AuthorizedRoles {
		err := a.db.AddResourceRole(ctx.Request().Context(), db.AddResourceRoleParams{
			ResourceUuid: resource.ResourceUuid,
			RoleUuid:     role,
		})
		if err != nil {
			return err
		}
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "resource created",
	})
	return nil
}

type addResourceRoleModel struct {
	ResourceUUID uuid.UUID   `json:"resourceUUID"`
	Roles        []uuid.UUID `json:"roles"`
}

func (a *API) AddResourceRole(ctx echo.Context) error {
	body, err := parseBodyToStruct[addResourceRoleModel](ctx.Request().Body)
	if err != nil {
		return err
	}
	for _, role := range body.Roles {
		err := a.db.AddResourceRole(ctx.Request().Context(), db.AddResourceRoleParams{
			ResourceUuid: body.ResourceUUID,
			RoleUuid:     role,
		})
		if err != nil {
			return err
		}
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "resource modified",
	})
	return nil
}
