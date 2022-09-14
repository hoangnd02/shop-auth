package main

import (
	"log"

	"github.com/hoanggggg5/shop-pkg/infrastructure/database"
	"github.com/hoanggggg5/shop/config"
	"github.com/hoanggggg5/shop/pkg/session"
	"github.com/hoanggggg5/shop/router"
)

func main() {
	db, err := database.New(&database.Config{
		Host:     config.Env.DatabaseHost,
		Port:     config.Env.DatabasePort,
		User:     config.Env.DatabaseUser,
		Password: config.Env.DatabasePass,
		DBName:   config.Env.DatabaseName,
	})
	if err != nil {
		log.Println(err)
	}

	session := session.InitSessionStore()

	app := router.InitializeRoutes(db, session)

	app.Listen(":3000")
}
