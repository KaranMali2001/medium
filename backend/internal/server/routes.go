package server

import (
	"medium/internal/middleware"

	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	e.GET("/blog", GetAllBlogs)
	e.POST("/blog", CreateNewBlog,middleware.VerifyToken)
	e.PUT("/blog/:id", UpdateBlog,middleware.VerifyToken)
	e.GET("/blog/:id", GetBlog)
	e.POST("/signUp", SignUp)
	e.POST("/signIn", SignIn)
}
