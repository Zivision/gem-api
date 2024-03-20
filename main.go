package main

import (
	"net/http"

	// Echo
	"github.com/labstack/echo/v4"
)

func main() {
	// Setup echo app
	e := echo.New();

	// Test end point ping
	e.GET("/ping", func (c echo.Context) error {
		return c.HTML(http.StatusOK, "<p>Pong!</p>");
	});

	// Add test data to map
	gems[diamond.ID] = diamond;

	// Gem api endpoints
	g := e.Group("/api/gems");
	g.GET("/", listGems);
	g.GET("/:id", listGem);

	// Start Server
	e.Logger.Fatal(e.Start(":8080"));
}
