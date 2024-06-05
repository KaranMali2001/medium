package database

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email    string `json:"email" gorm:"not null;uniqueIndex"`
	UserName string `json:"username" validate:"required" gorm:"not null;uniqueIndex"`
	Password string `json:"password" validate:"min=8,required"`
}
type Blog struct {
	gorm.Model

	Title     string `json:"title" gorm:"not null;unique"`
	Content   string `json:"content" gorm:"not null; unique"`
	Published bool   `json:"published"`
	UserEmail string `json:"user_id" gorm:"not null;index"`
	User      User   `gorm:"foreignKey:UserEmail;references:Email;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
}
type JWTClaims struct {
	Email string `json:"email"`

	jwt.RegisteredClaims
}
