package helpers

import (
	"github.com/creasty/defaults"
	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/shop/params"
)

func BodyParser[T any](c *fiber.Ctx, payload *T, prefix string) error {
	if err := c.BodyParser(payload); err != nil {
		return params.ErrServerInvalidBody
	}

	defaults.MustSet(payload)

	if err := Validate(payload, prefix); err != nil {
		return err
	} else {
		return nil
	}
}

func QueryParser[T any](c *fiber.Ctx, payload *T, prefix string) error {
	if err := c.QueryParser(payload); err != nil {
		return params.ErrServerInvalidQuery
	}

	defaults.MustSet(payload)

	if err := Validate(payload, prefix); err != nil {
		return err
	} else {
		return nil
	}
}
