package models

type Rank struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"string" gorm:"type:varchar(32)"`
	MinHours int    `json:"minHours"`
}
