package tokenprovider

import (
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}
type JwtClaims struct {
	jwt.RegisteredClaims
	UserClaims
}
