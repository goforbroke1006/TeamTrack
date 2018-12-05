package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Team struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Name       string    `gorm:"type:varchar(2048);unique_index"`
	ActiveFrom time.Time `gorm:"not null"`
	ActiveTill time.Time `gorm:"not null"`
}

type Member struct {
	gorm.Model

	Nick   string
	TeamID uint64
	Team   Team
}

type Location struct {
	gorm.Model

	MemberID  uint64
	Member    Member
	Latitude  float32
	Longitude float32
}
