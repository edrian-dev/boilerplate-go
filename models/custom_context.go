package models

import "github.com/labstack/echo"

// CustomContext ...
type CustomContext struct {
	echo.Context
	Body   echo.Map
	UserID string
}
