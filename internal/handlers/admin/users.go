package admin

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/shop-pkg/gpa"
	"github.com/hoanggggg5/shop-pkg/queries"
	"github.com/hoanggggg5/shop/internal/handlers/helpers"
	"github.com/hoanggggg5/shop/internal/handlers/resource/entities"
	"github.com/hoanggggg5/shop/internal/models"
)

func (h Handler) GetUsers(c *fiber.Ctx) error {
	type Params struct {
		UID   string           `query:"uid"`
		Email string           `query:"email"`
		Role  models.UserRole  `query:"role"`
		State models.UserState `query:"state"`
		queries.Period
		queries.Pagination
	}

	params := new(Params)

	if err := helpers.QueryParser(c, params, "admin.users"); err != nil {
		return err
	}

	ctx := context.Background()

	q := make([]gpa.Filter, 0)

	user, err := h.userUsecase.First(ctx, q...)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(entities.UserToEntity(user))
}
