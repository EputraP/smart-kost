package service

import (
	"os"
	"smart-kost-backend/dto"
	"smart-kost-backend/model"
	"smart-kost-backend/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(input dto.User) (string, error)
	 SignUp(input dto.User) (*dto.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

type UserServiceConfig struct {
	UserRepo repository.UserRepository
}

func NewUserService(config UserServiceConfig) UserService {
	return &userService{
		userRepo: config.UserRepo,
	}
}

func (ts *userService) Login(input dto.User) (string, error) {
	res, err := ts.userRepo.GetUserByUsername(&model.UserList{Username: input.UserName})
	if err != nil {
		return "", err
	}

	err=bcrypt.CompareHashAndPassword([]byte(res.Pass), []byte(input.Pass))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":input.UserName,
		"nbf":time.Now().Add(time.Hour*24*30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	
	return tokenString, nil
}

func (ts *userService) SignUp(input dto.User) (*dto.User, error) {
	hash, err :=bcrypt.GenerateFromPassword([]byte(input.Pass), 10)
	
	if err != nil {
		return nil, err
	}
	res, err := ts.userRepo.CreateUser(&model.UserList{Username: input.UserName, Pass: string(hash)})

	if err != nil {
		return nil, err
	}
	resp := &dto.User{
		UserName: res.Username,
		
	}

	return resp, err
}
