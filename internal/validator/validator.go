package valid

import (
	"github.com/alex-bogatiuk/wb_l0/internal/models"
	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

func ValidateOrderStruct(c *models.Order) error {

	err := Validator.Struct(c)

	if err != nil {
		return err
	}

	return nil
}
