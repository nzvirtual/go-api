package models

type Airport struct {
	ID   uint    `json:"id" gorm:"primaryKey"`
	Icao string  `json:"icao" gorm:"uniqueIndex;type:varchar(4)"`
	Name string  `json:"name" gorm:"type:varchar(50)"`
	Lat  float32 `json:"lat" gorm:"type:float(10,8)"`
	Lon  float32 `json:"lon" gorm:"type:float(10,7)"`
}
