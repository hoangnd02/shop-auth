package helpers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	pkg "github.com/hoanggggg5/shop-pkg/errors"
)

func Validate(value interface{}, prefix string) *pkg.Error {
	v := validate.Struct(value)

	v.AddTranslates(map[string]string{
		"Topic":    "topic",
		"State":    "state",
		"Role":     "role",
		"Type":     "type",
		"Limit":    "limit",
		"Page":     "page",
		"TimeFrom": "time_from",
		"TimeTo":   "time_to",
		"OrderBy":  "order_by",
		"Ordering": "order_ring",
	})

	e := make([]string, 0)

	if !v.Validate() {
		for _, errs := range v.Errors.All() {
			for _, err := range errs {
				e = append(e, fmt.Sprintf("%s.%s", prefix, err))
			}
		}
	}

	if len(e) > 0 {
		return pkg.NewError(fiber.StatusUnprocessableEntity, e...)
	} else {
		return nil
	}
}
