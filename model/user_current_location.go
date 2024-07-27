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
	IsSOS                   string    `json:"is_sos" gorm:"type:string;default:0"`
	IconColor               string    `json:"icon_color" gorm:"type:string;default:white"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

type GetUserCurrentLocation struct {
	Username   string `json:"username" gorm:"type:string"`
	IsOnline   string `json:"is_online" gorm:"type:string"`
	IsSOS      string `json:"is_sos" gorm:"type:string"`
	StatusName string `json:"status_name" gorm:"type:string"`
	Lat        string `json:"lat" gorm:"type:string"`
	Long       string `json:"long" gorm:"type:string"`
	IconColor  string `json:"icon_color" gorm:"type:string"`
}
