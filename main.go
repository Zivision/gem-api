package main

import (

	// Echo Lib
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Setup echo app
	e := echo.New();

	// Test end point PING
	e.GET("/ping", func (c echo.Context) error {
		return c.HTML(http.StatusOK, "<p>Pong!</p>");
	});

	// Start Server
	e.Logger.Fatal(e.Start(":8080"));
}
