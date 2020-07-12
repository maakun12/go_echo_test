package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Root Req")
}

func GetJson(c echo.Context) error {
	res := map[string]string{"key1": "val1", "key2": "val2"}
	return c.JSON(http.StatusOK, res)
}
