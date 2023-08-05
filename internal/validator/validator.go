package valid

import (
	"github.com/alex-bogatiuk/wb_l0/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/slog"
)

var vld = validator.New()

func ValidateOrderStruct(c *models.Order) error {
	err := vld.Struct(c)
	if err != nil {
		slog.Error("struct is not valid:", err)
		return err
	}

	return nil
}
