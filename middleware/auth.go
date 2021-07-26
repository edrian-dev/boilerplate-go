package middleware

import (
	"github.com/labstack/echo"
	_models "github.com/siends/siends-api/models"
)

// Auth ...
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customContext := &_models.CustomContext{}
		return next(customContext)
	}
}
