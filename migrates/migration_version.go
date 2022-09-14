package migrates

import "github.com/go-gormigrate/gormigrate/v2"

var ModelSchemaList = []*gormigrate.Migration{
	&initDatabase,
}
