package dto

type User struct {
	UserName                string               `json:"username" gorm:"type:string;not null"`
	Pass                    string               `json:"pass" gorm:"type:string;not null"`
	UserCurrentLocationData *UserCurrentLocation `json:"user_current_location_data" `
}

type UserResponse struct {
	UserId   int    `json:"user_id" gorm:"type:int;primaryKey"`
	Username string `json:"username" gorm:"type:string;not null"`
	IsOnline string `json:"is_online" gorm:"type:string;not null"`
	IsSOS    string `json:"is_sos" gorm:"type:string;not null"`
}

type UpdateUserOnlineSOS struct {
	UserId   int    `json:"user_id" `
	IsOnline string `json:"is_online" gorm:"type:string"`
	IsSOS    string `json:"is_sos" gorm:"type:string"`
}
