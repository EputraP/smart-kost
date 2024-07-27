package model

import (
	"time"
)

type UserCurrentLocation struct {
	UserCurrentLocationId   int       `json:"id" gorm:"type:int;primaryKey"`
	UserId                  int       `json:"user_id" gorm:"type:int;not null"`
	StatusId                int       `json:"status_id" gorm:"type:int;not null"`
	UserCurrentLocationLat  string    `json:"lat" gorm:"type:string;default:null"`
	UserCurrentLocationLong string    `json:"long" gorm:"type:string;default:null"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}
