package model

import (
	"time"
)

type UserList struct {
	UserId    int       `json:"user_id" gorm:"type:int;primaryKey"`
	Username  string    `json:"username" gorm:"type:string;not null"`
	Pass      string    `json:"pass" gorm:"type:string;not null"`
	IsOnline  string    `json:"is_online" gorm:"type:string;not null"`
	IsSOS     string    `json:"is_sos" gorm:"type:string;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
