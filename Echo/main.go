package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// e.Server.Addr = ":8080"

	// e.Use(middleware.Logger())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${latency_human}\n",
	}))

	e.GET("/", func(ctx echo.Context) error {
		// return ctx.String(http.StatusOK, "Hello, World!")
		return ctx.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})

	e.GET("/saludo", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Saludos!")
	})

	e.Start(":8080")
}
