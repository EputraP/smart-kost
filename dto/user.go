package dto

type User struct {
	UserName  string         `json:"username" gorm:"type:string;not null"`
	Pass  string             `json:"pass" gorm:"type:string;not null"`
}
