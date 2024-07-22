package dto

type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Pass     string `json:"pass" binding:"required"`
}

type LoginResponse struct {
	AccesToken   string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccesToken string `json:"access_token"`
}
