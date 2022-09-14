package identity

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/shop-pkg/gpa/filters"
	"github.com/hoanggggg5/shop/internal/handlers/helpers"
	"github.com/hoanggggg5/shop/internal/handlers/resource/entities"
	"github.com/hoanggggg5/shop/params"
	"github.com/hoanggggg5/shop/services"
)

const (
	passwordWrongErr = "password is incorrect"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	type Payload struct {
		Email    string `json:"email" validate:"required|email"`
		Password string `json:"password" validate:"required"`
	}

	payload := new(Payload)
	if err := helpers.BodyParser(c, payload, "identity.session"); err != nil {
		return err
	}

	ctx := c.Context()

	user, err := h.userUsecase.First(ctx, filters.WithFieldEqual("email", payload.Email))
	if err != nil {
		return err
	}

	if result := services.CheckPasswordHash(payload.Password, user.Password); !result {
		return c.Status(422).JSON(passwordWrongErr)
	}

	if user.Role == "" {
		return c.Status(422).JSON(params.ErrServerInternal)
	}

	if user.State == "Delete" {
		return c.Status(422).JSON(params.ErrServerInternal)
	}

	if user.State == "Banned" {
		return c.Status(422).JSON(params.ErrServerInternal)
	}

	session, err := h.sessionStore.Get(c)
	if err != nil {
		log.Println(err)
		return err
	}

	session.Set(user.UID, true)
	if err := session.Save(); err != nil {
		panic(err)
	}

	return c.JSON(entities.UserToEntity(user))
}

func (h *Handler) Logout(c *fiber.Ctx) error {
	session, err := h.sessionStore.Get(c)

	if err != nil {
		return c.Status(500).JSON(params.ErrServerInternal)
	}

	session.Destroy()
	session.Save()

	return c.JSON(200)
}
