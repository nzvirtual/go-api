package models

import "time"

type Role struct {
	ID        uint      `json:"-" gorm:"primarykey"`
	UserId    uint      `json:"userid" gorm:"index:userrole"`
	Name      string    `json:"role" gorm:"index:userrole"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
