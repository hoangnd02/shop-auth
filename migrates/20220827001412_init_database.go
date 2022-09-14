package migrates

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/volatiletech/null/v9"
	"gorm.io/gorm"
)

var initDatabase = gormigrate.Migration{
	ID: "20220827001412",
	Migrate: func(db *gorm.DB) error {
		type User struct {
			ID        int64       `gorm:"primaryKey;autoIncrement;not null"`
			UID       string      `gorm:"type:character varying;not null;uniqueIndex:index_users_on_uid"`
			Username  null.String `gorm:"type:character varying;uniqueIndex:index_users_on_username"`
			Email     string      `gorm:"type:character varying;not null;uniqueIndex:index_users_on_email"`
			Password  string      `gorm:"type:character varying;not null"`
			Role      string      `gorm:"type:character varying;not null;default:member"`
			State     string      `gorm:"type:character varying;not null;default:pending"`
			CreatedAt time.Time   `gorm:"type:timestamp;not null"`
			UpdatedAt time.Time   `gorm:"type:timestamp;not null"`
		}

		return db.AutoMigrate(
			User{},
		)
	},
	Rollback: func(db *gorm.DB) error {
		return db.Migrator().DropTable(
			"users",
		)
	},
}
