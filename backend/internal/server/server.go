package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetAllBlogs(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "returning all the blogs")
}
func CreateNewBlog(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "new blog created successfully")
}
func UpdateBlog(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "blog updated successfully")
}
func GetBlog(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "returning particular blog")
}
func SignIn(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "sign in successfully")
}
func SignUp(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "sign up successfully")
}
