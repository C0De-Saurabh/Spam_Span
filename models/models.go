package models

import (
	"time"
)

// User model
type User struct {
	ID                 uint   `gorm:"primaryKey"`
	Name               string `gorm:"not null"`
	Phone              string `gorm:"unique;not null"`
	Password           string `gorm:"not null"`
	Email              string
	Contacts           []Contact `gorm:"foreignKey:UserID"` // One-to-many relationship
	SpamReportsToday   int       `gorm:"default:0"`
	LastSpamReportDate string    `gorm:"default:null"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

// Contact model
type Contact struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"` // Foreign key for User
	Name      string `gorm:"not null"` // Contact name
	Phone     string `gorm:"not null"` // Contact phone number
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GlobalContact model
type GlobalContact struct {
	ID           uint   `gorm:"primaryKey"`
	Phone        string `gorm:"unique;not null"`
	Name         string `gorm:"default:'Unknown'"`
	SpamReported int    `gorm:"default:0"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
