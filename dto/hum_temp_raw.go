package dto

type CreateHumTempRaw struct {
	SensorID int    `json:"sensor_id" gorm:"type:int;not null"`
	Hum      string `json:"hum" gorm:"type:varchar"`
	Tem      string `json:"tem" gorm:"type:varchar"`
}
