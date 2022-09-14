package identity

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/shop/internal/handlers/helpers"
	"github.com/hoanggggg5/shop/internal/models"
	"github.com/hoanggggg5/shop/services"
	"github.com/hoanggggg5/shop/utils"
)

func (h *Handler) Register(c *fiber.Ctx) error {
	type Payload struct {
		Email    string `json:"email" validate:"required|email"`
		Password string `json:"password" validate:"required"`
	}

	payload := new(Payload)
	if err := helpers.BodyParser(c, payload, "identity.user"); err != nil {
		return err
	}

	ctx := c.Context()

	user := &models.User{
		UID:      utils.RandomUID(),
		Email:    payload.Email,
		Password: services.HashPassword(payload.Password),
		State:    "Pending",
		Role:     "Member",
	}

	if err := h.userUsecase.Create(ctx, user); err != nil {
		return err
	}

	session, err := h.sessionStore.Get(c)
	if err != nil {
		return err
	}

	session.Set("uid", user.UID)
	if err := session.Save(); err != nil {
		panic(err)
	}

	return c.SendStatus(200)
}
