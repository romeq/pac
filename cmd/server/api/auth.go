package api

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

var (
	errMissingAuthentication     = fmt.Errorf("missing authentication")
	errWrongAuthenticationMethod = fmt.Errorf("wrong authentication method")
	errAuthenticationFailed      = fmt.Errorf("authentication failed")
)

func (a *API) AdminAuthentication(password string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(auth) < 2 {
				return errMissingAuthentication
			} else if strings.TrimSpace(auth[0]) != "Basic" {
				return errWrongAuthenticationMethod
			}

			userPassword := strings.TrimSpace(auth[1])
			decoded, err := base64.RawStdEncoding.DecodeString(userPassword)
			if err != nil {
				return err
			}

			if strings.TrimSpace(string(decoded))[1:] != password {
				return errAuthenticationFailed
			}

			return next(c)
		}
	}
}
