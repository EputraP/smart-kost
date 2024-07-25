package dto

type UserCurrentLocation struct {
	UserCurrentLocationId   int    `json:"id" gorm:"type:int;primaryKey"`
	UserId                  int    `json:"user_id" gorm:"type:int;not null"`
	StatusId                int    `json:"status_id" gorm:"type:int;not null"`
	UserCurrentLocationLat  string `json:"lat" gorm:"type:string"`
	UserCurrentLocationLong string `json:"long" gorm:"type:string"`
	CurrentLocation         string `json:"current_location" gorm:"type:string"`
}
