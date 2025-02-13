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
	Village        string `json:"village" gorm:"type:string"`
	Subdistrict string `json:"subdistrict" gorm:"type:string"`
	City        string `json:"city" gorm:"type:string"`
	Province    string `json:"state" gorm:"type:string"`
}

type GetUserCurrentLocationResponse struct {
	UserId   int  `json:"user_id" gorm:"type:integer"`
	Username   string  `json:"username" gorm:"type:string"`
	IsOnline   string  `json:"is_online" gorm:"type:string"`
	IsSOS      string  `json:"is_sos" gorm:"type:string"`
	StatusName string  `json:"status_name" gorm:"type:string"`
	Lat        float64 `json:"lat" gorm:"type:float"`
	Long       float64 `json:"long" gorm:"type:float"`
	Address    string  `json:"address" gorm:"type:string"`
	IconColor  string  `json:"icon_color" gorm:"type:string"`
}

type UpdateUserSOS struct {
	UserId int    `json:"user_id" `
	IsSOS  string `json:"is_sos" gorm:"type:string"`
}

type IconColorRGB struct {
	Red   int `json:"red" gorm:"type:integer"`
	Green int `json:"green" gorm:"type:integer"`
	Blue  int `json:"blue" gorm:"type:integer"`
}

type GetSingleUserCurrentLocationResponse struct {
	UserId   int  `json:"user_id" gorm:"type:integer"`
	Username   string  `json:"username" gorm:"type:string"`
	IsOnline   string  `json:"is_online" gorm:"type:string"`
	IsSOS      string  `json:"is_sos" gorm:"type:string"`
	StatusName string  `json:"status_name" gorm:"type:string"`
	Lat        float64 `json:"lat" gorm:"type:float"`
	Long       float64 `json:"long" gorm:"type:float"`
	DisplayName    string  `json:"display_name" gorm:"type:string"`
	Address    Address  `json:"address"`
	LastOnline    string  `json:"last_online" gorm:"type:string"`
	LastLocationData    string  `json:"last_location_data" gorm:"type:string"`
	IconColor  string  `json:"icon_color" gorm:"type:string"`
}
