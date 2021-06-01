package models

type Airport struct {
	ID   uint    `json:"id" gorm:"primaryKey"`
	Icao string  `json:"icao" gorm:"uniqueIndex;size:4"`
	Name string  `json:"name"`
	Lat  float32 `json:"lat" gorm:"float(10,8)"`
	Lon  float32 `json:"lon" gorm:"float(10,7)"`
}
