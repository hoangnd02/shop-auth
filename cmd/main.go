package main

import (
	"log"
	"os"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/hoanggggg5/shop-pkg/infrastructure/database"
	"github.com/hoanggggg5/shop/config"
	"github.com/hoanggggg5/shop/migrates"
	"github.com/hoanggggg5/shop/pkg/session"
	"github.com/hoanggggg5/shop/router"
	"github.com/mkideal/cli"
)

// root command
type rootT struct {
	cli.Helper
}

var root = &cli.Command{
	Desc: "this is root command",
	Argv: func() interface{} { return new(rootT) },
}

func main() {
	if err := config.Initialize(); err != nil {
		panic(err)
	}

	if err := cli.Root(root,
		cli.Tree(api),
		cli.Tree(migration),
	).Run(os.Args[1:]); err != nil {
		log.Panicln(err)
		os.Exit(1)
	}
}

var api = &cli.Command{
	Name: "api",
	Desc: "This command will run auth api",
	Fn: func(ctx *cli.Context) error {
		if err := config.Initialize(); err != nil {
			panic(err)
		}

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

		return nil
	},
}

var migration = &cli.Command{
	Name: "migration",
	Desc: "this is migration command",
	Fn: func(ctx *cli.Context) error {
		db, err := database.New(&database.Config{
			Host:     config.Env.DatabaseHost,
			Port:     config.Env.DatabasePort,
			User:     config.Env.DatabaseUser,
			Password: config.Env.DatabasePass,
			DBName:   config.Env.DatabaseName,
		})
		if err != nil {
			return err
		}

		migrate := gormigrate.New(db, gormigrate.DefaultOptions, migrates.ModelSchemaList)

		return migrate.Migrate()
	},
}
