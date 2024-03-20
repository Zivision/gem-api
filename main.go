package main

import (
	// Echo
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Setup echo app
	e := echo.New();

	// Middleware
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Gem api endpoints
	g := e.Group("/api/gems");
	g.GET("/", listGems);
	g.GET("/:id", listGem);
	g.POST("/", createGem);
	g.PUT("/:id", updateGem);

	// Start Server
	e.Logger.Fatal(e.Start(":8080"));
}
