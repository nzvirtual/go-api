package models

import (
	"time"
)

type User struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	Email             string    `json:"email" gorm:"not null"`
	Firstname         string    `json:"firstname" gorm:"not null"`
	Lastname          string    `json:"lastname" gorm:"not null"`
	Password          string    `json:"-"`
	Verified          bool      `json:"verified" gorm:"default:false"`
	VerificationToken uint      `json:"-" gorm:"size:6"`
	LastAirportID     uint      `json:"-"`
	LastAirport       Airport   `json:"lastAirport"`
	Hours             float32   `json:"hours" gorm:"float(10,2)"`
	RankID            uint      `json:"-"`
	Rank              Rank      `json:"rank"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
