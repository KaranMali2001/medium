package database

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `json:"id" `
	Email    string `json:"email" gorm:"not null;uniqueIndex"`
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"min=8,required"`
}
type Blog struct {
	gorm.Model
	ID        uint   `json:"id"`
	Title     string `json:"title" gorm:"not null;unique"`
	Content   string `json:"content" gorm:"not null; unique"`
	Published bool   `json:"published"`
	UserID    uint   `json:"user_id" gorm:"not null;index"`
	User      User   `gorm:"foreignKey:UserID;references:ID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
}
type JWTClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
