package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"

	"urban-fork/render"
)

func main() {
	fmt.Println("Hello world")

	e := echo.New()

	// Little bit of middlewares for housekeeping
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	// This will initiate our template renderer
	render.NewTemplateRenderer(e, "public/*.html")
	e.GET("/hello", func(e echo.Context) error {
		return e.Render(http.StatusOK, "index", nil)
	})

	e.Logger.Fatal(e.Start(":4040"))
}
