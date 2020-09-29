package main

import (
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Setup proxy
	url1, err := url.Parse("https://859c9b77b8a1.ngrok.io")
	if err != nil {
		e.Logger.Fatal(err)
	}
	url2, err := url.Parse("https://859c9b77b8a1.ngrok.io")
	if err != nil {
		e.Logger.Fatal(err)
	}
	targets := []*middleware.ProxyTarget{
		{
			URL: url1,
		},
		{
			URL: url2,
		},
	}
	e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

	e.Logger.Fatal(e.Start(":8080"))
}
