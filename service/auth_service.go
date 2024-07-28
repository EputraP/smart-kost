package service

import (
	"encoding/json"
	"math/rand"
	"smart-kost-backend/dto"
	"smart-kost-backend/errs"
	"smart-kost-backend/model"
	"smart-kost-backend/repository"
	"smart-kost-backend/util/hasher"
	"smart-kost-backend/util/tokenprovider"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(input dto.LoginBody) (*dto.LoginResponse, error)
	SignUp(input dto.LoginBody) (*dto.User, error)
	Logout(input int) (string, error)
}

type authService struct {
	userRepo                repository.UserRepository
	userCurrentLocationRepo repository.UserCurrentLocationRepo
	hasher                  hasher.Hasher
	jwtProvider             tokenprovider.JWTTokenProvider
}

type AuthServiceConfig struct {
	UserRepo                repository.UserRepository
	UserCurrentLocationRepo repository.UserCurrentLocationRepo
	Hasher                  hasher.Hasher
	JwtProvider             tokenprovider.JWTTokenProvider
}

func NewAuthService(config AuthServiceConfig) AuthService {
	return &authService{
		userRepo:                config.UserRepo,
		userCurrentLocationRepo: config.UserCurrentLocationRepo,
		hasher:                  config.Hasher,
		jwtProvider:             config.JwtProvider,
	}
}

func (ts authService) Login(input dto.LoginBody) (*dto.LoginResponse, error) {
	lowerUsername := strings.ToLower(input.Username)
	account, err := ts.userRepo.GetUserByUsername(&model.UserList{Username: lowerUsername})
	if err != nil {
		return nil, err
	}
	_, err = ts.userRepo.UpdateIsOnline(&model.UserList{UserId: account.UserId, IsOnline: "1"})
	if err != nil {
		return nil, err
	}

	passwordOk, err := ts.hasher.IsEqual(account.Pass, input.Pass)

	if err != nil {
		return nil, err
	}

	if !passwordOk {
		return nil, errs.PasswordDoesntMatch
	}

	userClaims := model.UserList{}

	if account != nil {
		userClaims.UserId = account.UserId
		userClaims.Username = account.Username
	}

	return ts.generateLoginResponse(&userClaims)

}

func (ts authService) SignUp(input dto.LoginBody) (*dto.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Pass), 10)
	usernameLower := strings.ToLower(input.Username)

	if err != nil {
		return nil, err
	}
	res, err := ts.userRepo.CreateUser(&model.UserList{Username: usernameLower, Pass: string(hash)})

	if err != nil {
		return nil, err
	}

	randomColor := GenerateRandomColor()
	stringColor, err := json.Marshal(randomColor)
	if err != nil {
		return nil, err
	}

	resUserCurrentLocation, err := ts.userCurrentLocationRepo.CreateCurrentUserData(&model.UserCurrentLocation{
		UserId:    res.UserId,
		StatusId:  2,
		IconColor: string(stringColor),
	})
	if err != nil {
		return nil, err
	}

	userCurrentLocationData := &dto.UserCurrentLocation{
		UserCurrentLocationId:   resUserCurrentLocation.UserCurrentLocationId,
		UserId:                  resUserCurrentLocation.UserId,
		StatusId:                resUserCurrentLocation.StatusId,
		UserCurrentLocationLat:  resUserCurrentLocation.UserCurrentLocationLat,
		UserCurrentLocationLong: resUserCurrentLocation.UserCurrentLocationLong,
	}

	resp := &dto.User{
		UserName:                res.Username,
		UserCurrentLocationData: userCurrentLocationData,
	}

	return resp, err
}

func (ts authService) Logout(input int) (string, error) {

	_, err := ts.userRepo.UpdateIsOnline(&model.UserList{UserId: input, IsOnline: "0"})
	if err != nil {
		return "", err
	}

	return "", err
}

func GenerateRandomColor() *dto.IconColorRGB {
	var res *dto.IconColorRGB
	red := rand.Intn(255)
	green := rand.Intn(255)
	blue := rand.Intn(255)
	res = &dto.IconColorRGB{
		Red:   red,
		Blue:  blue,
		Green: green,
	}
	return res
}

func (ts authService) generateLoginResponse(user *model.UserList) (*dto.LoginResponse, error) {
	accesToken, err := ts.jwtProvider.GenerateAccessToken(*user)

	if err != nil {
		return nil, err
	}

	refreshToken, err := ts.jwtProvider.GenerateRefreshToken(*user)

	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccesToken:   accesToken,
		RefreshToken: refreshToken,
	}, nil
}
