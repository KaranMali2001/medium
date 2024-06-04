package middleware

import (
	"log"

	"github.com/go-playground/validator/v10"
)

func ValidateReq(reqBody interface{}) error {

	v := validator.New()
	entity := reqBody
	err := v.Struct(entity)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
