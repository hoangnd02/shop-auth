package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/hoanggggg5/shop-pkg/log"
	"github.com/hoanggggg5/shop/types"
)

var (
	Env types.ENV
)

func Initialize() error {
	if err := env.Parse(&Env); err != nil {
		return err
	}

	log.New(Env.ApplicationName)

	return nil
}
