package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/shop-pkg/log"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	if err != nil {
		return c.Status(code).JSON(err.Error())
	}

	return nil
}

func DefaultStackTraceHandler(_ *fiber.Ctx, e interface{}) {
	log.Errorf("Panic: %v\n", e)
}
