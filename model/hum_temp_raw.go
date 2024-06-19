package model

import (
	"time"

	"gorm.io/gorm"
)

type HumTemRaw struct {
	DataID    int            `json:"data_id" gorm:"type:int;primaryKey"`
	SensorID  int            `json:"sensor_id" gorm:"type:int;not null"`
	Hum       string         `json:"hum" gorm:"type:varchar"`
	Tem       string         `json:"tem" gorm:"type:varchar"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
