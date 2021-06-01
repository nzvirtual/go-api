package models

type Rank struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"string"`
	MinHours int    `json:"minHours"`
}
