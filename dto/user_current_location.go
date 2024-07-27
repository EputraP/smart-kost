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

type GetUserCurrentLocationResponse struct {
	Username   string `json:"username" gorm:"type:string"`
	IsOnline   string `json:"is_online" gorm:"type:string"`
	IsSOS      string `json:"is_sos" gorm:"type:string"`
	StatusName string `json:"status_name" gorm:"type:string"`
	Lat        string `json:"lat" gorm:"type:string"`
	Long       string `json:"long" gorm:"type:string"`
	Address    string `json:"address" gorm:"type:string"`
}
