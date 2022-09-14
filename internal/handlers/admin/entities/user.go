package entities

import (
	"time"

	"github.com/hoanggggg5/shop/internal/models"
	"github.com/volatiletech/null/v9"
)

type User struct {
	ID        int64            `json:"id,omitempty"`
	UID       string           `json:"uid,omitempty"`
	Username  null.String      `json:"username,omitempty"`
	Email     string           `json:"email,omitempty"`
	Role      models.UserRole  `json:"role,omitempty"`
	State     models.UserState `json:"state,omitempty"`
	Data      null.String      `json:"data,omitempty"`
	CreatedAt time.Time        `json:"created_at,omitempty"`
	UpdatedAt time.Time        `json:"updated_at,omitempty"`
}

func UserToEntity(user *models.User) *User {
	return &User{
		ID:        user.ID,
		UID:       user.UID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      models.UserRole(user.Role),
		State:     models.UserState(user.State),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
