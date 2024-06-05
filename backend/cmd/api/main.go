package main

import (

	"medium/internal/server"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	server.Routes(e)
	
	e.Start(":8080")
}
