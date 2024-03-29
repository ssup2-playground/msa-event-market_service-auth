package entity

import (
	"time"

	"gorm.io/gorm"

	"github.com/ssup2-playground/msa-event-market_service-auth/pkg/entity/uuid"
)

type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
)

func IsValidUserRole(role string) bool {
	if role == string(UserRoleAdmin) {
		return true
	} else if role == string(UserRoleUser) {
		return true
	}
	return false
}

type UserInfo struct {
	ID        uuid.EntityUUID `gorm:"primaryKey;type:binary(16)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	LoginID string   `gorm:"unique;size:20"` // Unique key
	Role    UserRole `gorm:"size:20"`
	Phone   string   `gorm:"size:13"`
	Email   string   `gorm:"size:40"`
}
