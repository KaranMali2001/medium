package middleware

import (
	"log"
	"medium/internal/database"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
)

var JWTKey = []byte("my_key")

func CreateToken(email string, id uint) (string, error) {
	claims := database.JWTClaims{

		Email:            email,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTKey)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenString, nil
}
func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		auth := ctx.Request().Header.Get("authorization")
		if auth == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "jwt is empty")
		}
		if strings.HasPrefix(auth, "Bearer ") {
			auth = strings.TrimPrefix(auth, "Bearer ")
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "jwt does not have bearer")
		}

		token, err := jwt.ParseWithClaims(auth, &database.JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(JWTKey), nil
		})
		if err != nil {
			log.Println(err)
			return ctx.JSON(http.StatusInternalServerError, "validating token error")
		}
		claims, ok := token.Claims.(*database.JWTClaims)
		if !ok || !token.Valid {
			log.Println(ok, token.Valid)
			return ctx.JSON(http.StatusInternalServerError, "claims error")
		}
		ctx.Set("claims", claims)

		return next(ctx)
	}
}
