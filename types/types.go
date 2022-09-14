package types

type ENV struct {
	Domain          string `env:"DOMAIN" envDefault:""`
	ApplicationName string `env:"APP_NAME" envDefault:"auth"`

	DatabaseHost string `env:"DATABASE_HOST" envDefault:"localhost"`
	DatabasePort int    `env:"DATABASE_PORT" envDefault:"5432"`
	DatabaseUser string `env:"DATABASE_USER" envDefault:"root"`
	DatabasePass string `env:"DATABASE_PASS" envDefault:"123456"`
	DatabaseName string `env:"DATABASE_NAME" envDefault:"auth"`
}
