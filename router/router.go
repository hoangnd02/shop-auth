package router

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/hoanggggg5/shop-pkg/gpa/filters"
	"github.com/hoanggggg5/shop-pkg/infrastructure/repository"
	"github.com/hoanggggg5/shop/config"
	"github.com/hoanggggg5/shop/internal/handlers/admin"
	"github.com/hoanggggg5/shop/internal/handlers/identity"
	"github.com/hoanggggg5/shop/internal/models"
	"github.com/hoanggggg5/shop/internal/usecases"
	"github.com/hoanggggg5/shop/params"
	"github.com/hoanggggg5/shop/router/middlewares"
	"github.com/hoanggggg5/shop/utils"
	"gorm.io/gorm"
)

func InitializeRoutes(
	db *gorm.DB,
	session *session.Store,
) *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit:               10 * 1024 * 1024, // this is the default limit of 10MB
		EnableTrustedProxyCheck: true,
		ProxyHeader:             "X-Forwarded-For",
		AppName:                 config.Env.ApplicationName,
		ErrorHandler:            middlewares.ErrorHandler,
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: middlewares.DefaultStackTraceHandler,
	}))

	userRepository := repository.New(db, models.User{})
	userUsecase := usecases.NewUserUsecase(userRepository)

	apiV2 := app.Group("/api/v2")

	apiV2.All("/auth/*", middlewares.MustAuth, func(c *fiber.Ctx) error {
		session, err := session.Get(c)

		if err != nil {
			return c.Status(500).JSON(params.ErrServerInternal)
		}

		uid := session.Get("uid")

		if uid == nil {
			return c.Status(500).JSON(params.ErrServerInternal)
		}

		user, err := userUsecase.First(context.Background(), filters.WithFieldEqual("uid", uid))
		if err != nil {
			return nil
		}

		if len(user.Email) == 0 {
			return c.Status(500).JSON(params.ErrServerInternal)
		}

		if user.State != "Active" {
			return c.Status(500).JSON(params.ErrServerInternal)
		}

		jwt_token, err := utils.GenerateJWT(user)

		if err != nil {
			return c.Status(500).JSON(params.ErrServerInternal)
		}

		jwt_token = "Bearer " + jwt_token
		c.Set("Authorization", jwt_token)

		return c.SendStatus(200)
	})

	admin.NewRouter(
		apiV2.Group("/admin"),
		userUsecase,
		session,
	)

	identity.NewRouter(
		apiV2.Group("/identity"),
		userUsecase,
		session,
	)

	return app
}
