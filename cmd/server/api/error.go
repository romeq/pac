package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
)

var (
	ErrDuplicate                  = fmt.Errorf("record already exists")
	ErrForeignKeyConstraintFailed = fmt.Errorf("a record could not be found from linked table (constraint failed)")
	ErrTablesMissing              = fmt.Errorf("database has not been setup correctly on the server")
	ErrUnknown                    = fmt.Errorf("unknown error occured")
)

func dberror(err *pgconn.PgError) error {
	switch err.Code {
	case "23505":
		return ErrDuplicate
	case "42P01":
		log.Println("CRITICAL: some critical tables are missing, please re-run migrations")
		return ErrTablesMissing
	case "23503":
		return ErrForeignKeyConstraintFailed
	default:
		log.Printf("ERROR (%s): %s\n", err.Code, err.Message)
		return ErrUnknown
	}
}

type errorResponse struct {
	Error string `json:"error"`
}

func (a *API) ErrorHandler(err error, ctx echo.Context) {
	var pgconnErr *pgconn.PgError
	if errors.As(err, &pgconnErr) {
		err = dberror(pgconnErr)
	}

	ctx.JSON(http.StatusInternalServerError, &errorResponse{
		Error: err.Error(),
	})
}
