package session

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

func InitSessionStore() *session.Store {
	storage := redis.New(redis.Config{
		Host:  "localhost",
		Port:  6379,
		Reset: false,
	})

	return session.New(session.Config{
		Storage:        storage,
		Expiration:     7 * time.Hour,
		CookiePath:     "/",
		CookieHTTPOnly: true,
	})
}
