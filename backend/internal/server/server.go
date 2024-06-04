package server

import (
	"log"
	"medium/internal/database"
	"medium/internal/middleware"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllBlogs(ctx echo.Context) error {
	db := database.Db
	var blogs []database.Blog
	result := db.Find(&blogs)
	if result.Error != nil {
		log.Println(result.Error)
		return ctx.JSON(http.StatusInternalServerError, "error while finding all post")
	}
	return ctx.JSON(http.StatusOK, blogs)
}
func CreateNewBlog(ctx echo.Context) error {
	db := database.Db
	claims := ctx.Get("claims").(*database.JWTClaims)
	blog := new(database.Blog)
	if err := ctx.Bind(&blog); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, "error while binding data")
	}
	err := middleware.ValidateReq(blog)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, "error while validating ")
	}
	result := db.Create(blog)
	if result.Error != nil {
		log.Println(result.Error)
		return ctx.JSON(http.StatusInternalServerError, "error while creating new blog")
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":    "new blog created successfully",
		"created by": claims.Email,
	})
}
func UpdateBlog(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "blog updated successfully")
}
func GetBlog(ctx echo.Context) error {
	id := ctx.Param("id")
	var blog database.Blog
	db := database.Db
	result := db.Where("id = ?", id).Model(&blog)
	if result.Error != nil {
		log.Println(result.Error)
		return ctx.JSON(http.StatusInternalServerError, "error while finding ")
	}

	return ctx.JSON(http.StatusOK, blog)
}
func SignIn(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	var user database.User
	db := database.Db
	result := db.Where("username = ?,password= ?", username, password).First(&user)
	if result.Error != nil {
		log.Println(result.Error)
		return ctx.JSON(http.StatusInternalServerError, "error while finding")
	}
	token, err := middleware.CreateToken(user.Email)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "sign in successfully",
		"token":   token,
	})
}
func SignUp(ctx echo.Context) error {
	user := new(database.User)
	db := database.Db
	if err := ctx.Bind(user); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, "error while binding")
	}
	err := middleware.ValidateReq(user)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, "error while validating ")
	}
	result := db.Create(user)
	if result.Error != nil {
		log.Println(result.Error)
		return ctx.JSON(http.StatusInternalServerError, "error while creating ")
	}
	token, err := middleware.CreateToken(user.Email)
	if err != nil {
		log.Println(err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "sign up successfully",
		"token ":  token,
	})
}
