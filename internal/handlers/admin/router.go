package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/hoanggggg5/shop/internal/usecases"
)

type Handler struct {
	userUsecase  usecases.UserUsecase
	sessionStore *session.Store
}

func NewRouter(
	router fiber.Router,
	userUsecase usecases.UserUsecase,
	session *session.Store,
) {
	handler := Handler{
		userUsecase:  userUsecase,
		sessionStore: session,
	}

	router.Get("/users", handler.GetUsers)
}
