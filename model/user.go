package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const USER_TYPE = "user"
const ADMIN_TYPE = "admin"
const ALL_TYPE = "all"

type UUIDPrimaryKey struct {
	ID        string `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *UUIDPrimaryKey) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}

type User struct {
	UUIDPrimaryKey
	Name     string  `gorm:"type:varchar(100)" json:"name"`
	Email    string  `gorm:"type:varchar(100)" json:"email"`
	Phone    string  `gorm:"type:varchar(20)" json:"phone"`
	Amount   float64 `gorm:"type:decimal(12)" json:"amount"`
	Password string  `gorm:"type:varchar" json:"password"`
	UserType string  `gorm:"type:varchar(100)" json:"user_type"`
	Pin      string  `gorm:"type:varchar" json:"pin"`
	Image    string  `gorm:"type:varchar" json:"image"`
	Address  string  `gorm:"type:text" json:"address"`
}
