package tokenprovider

import (
	"log"
	"smart-kost-backend/errs"
	"smart-kost-backend/model"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTTokenProvider interface {
	GenerateRefreshToken(user model.UserList) (string, error)
	GenerateAccessToken(user model.UserList) (string, error)
	ValidateToken(token string) (*JwtClaims, error)
}

type jwtTokenProvider struct {
	issuer               string
	secret               string
	refreshTokenDuration int
	accessTokenDuration  int
}

func NewJWT(issuer string, secret string, refreshTokenDuration int, accessTokenDuration int) JWTTokenProvider {
	return &jwtTokenProvider{
		issuer:               issuer,
		secret:               secret,
		refreshTokenDuration: refreshTokenDuration,
		accessTokenDuration:  accessTokenDuration,
	}
}

func (p *jwtTokenProvider) GenerateAccessToken(user model.UserList) (string, error) {
	return p.generateToken(user, time.Duration(p.accessTokenDuration)*time.Minute)
}

func (p *jwtTokenProvider) GenerateRefreshToken(user model.UserList) (string, error) {
	return p.generateToken(user, time.Duration(p.refreshTokenDuration)*time.Minute)
}

func (p *jwtTokenProvider) generateToken(user model.UserList, expiresIn time.Duration) (string, error) {
	claims := JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    p.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserClaims: UserClaims{
			UserID:   strconv.Itoa(user.UserId),
			Username: user.Username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(p.secret))
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenStr, nil
}

func (p *jwtTokenProvider) ValidateToken(token string) (*JwtClaims, error) {
	claims := JwtClaims{}

	jwtToken, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(p.secret), nil
	})

	println(claims.Issuer)
	println(p.issuer)
	if jwtToken == nil || !jwtToken.Valid {
		return nil, errs.InvalidToken
	}

	if err != nil {
		return nil, err
	}

	if claims.Issuer != p.issuer {
		return nil, errs.InvalidIssuer
	}

	return &claims, nil
}
