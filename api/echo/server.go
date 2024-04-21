package echoserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Start(port string) {
	r := echo.New()

	r.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "echo pong"})
	})

	if err := r.Start(port); err != nil {
		r.Logger.Fatal(err)
	}
}
