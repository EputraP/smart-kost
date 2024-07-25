package dto

type UserCurrentLocation struct {
	UserCurrentLocationId   int    `json:"id" gorm:"type:int;primaryKey"`
	UserId                  int    `json:"user_id" gorm:"type:int;not null"`
	StatusId                int    `json:"status_id" gorm:"type:int;not null"`
	UserCurrentLocationLat  string `json:"lat" gorm:"type:string"`
	UserCurrentLocationLong string `json:"long" gorm:"type:string"`
}

type UpdateUserLocation struct {
	UserId                  int    `json:"user_id" gorm:"type:int;not null"`
	UserCurrentLocationLat  string `json:"lat" gorm:"type:string"`
	UserCurrentLocationLong string `json:"long" gorm:"type:string"`
}

type GetLocation struct {
	DisplayName string  `json:"display_name" gorm:"type:string"`
	Address     Address `json:"address" `
}
type Address struct {
	Road        string `json:"road" gorm:"type:string"`
	Subdistrict string `json:"subdistrict" gorm:"type:string"`
	City        string `json:"city" gorm:"type:string"`
	Province    string `json:"province" gorm:"type:string"`
}
