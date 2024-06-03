package main

import (
	"medium/internal/database"
	"medium/internal/server"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	server.Routes(e)
	database.Start()
	e.Start(":8080")
}
