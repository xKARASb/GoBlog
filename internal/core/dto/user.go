package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserDB struct {
	UserId                 uuid.UUID `json:"user_id" db:"user_id"`
	Email                  string    `json:"email" db:"email"`
	PasswordHash           string    `db:"password_hash"`
	Role                   string    `json:"role" db:"role"`
	RefreshToken           string    `json:"refresh_token" db:"refresh_token"`
	RefreshTokenExpiryTime time.Time `db:"refresh_token_expiry_time"`
}
