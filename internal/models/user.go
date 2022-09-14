package models

import (
	"time"

	"github.com/volatiletech/null/v9"
)

type UserState string

const (
	UserStateActive  UserState = "active"
	UserStatePending UserState = "pending"
	UserStateDeleted UserState = "deleted"
	UserStateBanned  UserState = "banned"
	UserStateLocked  UserState = "locked"
)

var UserStates = []UserState{
	UserStateActive,
	UserStatePending,
	UserStateDeleted,
	UserStateBanned,
	UserStateLocked,
}

type UserRole string

const (
	UserRoleMember     UserRole = "member"
	UserRoleAdmin      UserRole = "admin"
	UserRoleSuperAdmin UserRole = "superadmin"
)

var UserRoles = []UserRole{
	UserRoleMember,
	UserRoleAdmin,
	UserRoleSuperAdmin,
}

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

func (u User) TableName() string {
	return "users"
}
