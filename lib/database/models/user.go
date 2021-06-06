package models

import (
	"strings"
	"time"
)

type User struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	Email             string    `json:"email" gorm:"type:varchar(155);uniqueIndex"`
	Firstname         string    `json:"firstname" gorm:"type:varchar(155)"`
	Lastname          string    `json:"lastname" gorm:"type:varchar(155)"`
	Password          string    `json:"-" gorm:"type:varchar(128)"`
	Verified          bool      `json:"verified" gorm:"default:false"`
	VerificationToken string    `json:"-" gorm:"type:varchar(15)"`
	LastAirportID     uint      `json:"-"`
	LastAirport       Airport   `json:"lastAirport"`
	Hours             float32   `json:"hours" gorm:"float(10,2)"`
	RankID            uint      `json:"-"`
	Rank              Rank      `json:"rank"`
	Roles             []Role    `json:"roles"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

func (u *User) HasRole(role string) bool {
	for i := 0; i < len(u.Roles); i++ {
		if strings.ToLower(u.Roles[i].Name) == strings.ToLower(role) {
			return true
		}
	}

	return false
}
