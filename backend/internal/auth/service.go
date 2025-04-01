package auth

import (
	"errors"

	"github.com/Vitaljano/invly/backend/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepo *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepo,
	}
}

func (service AuthService) Login(email, password string) (string, error) {
	user, err := service.UserRepository.FindByEmail(email)

	if err != nil {
		return "", errors.New(ErrorWrongCredentials)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", errors.New(ErrorWrongCredentials)
	}

	return user.Email, nil

}

func (service AuthService) Register(email, name, password string) (string, error) {

	existedUser, _ := service.UserRepository.FindByEmail(email)

	if existedUser != nil {
		return "", errors.New(ErrorUserExists)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	user := &user.User{
		Email:    email,
		Password: string(passwordHash),
		Name:     name,
	}

	user, err = service.UserRepository.Create(user)

	if err != nil {
		return "", err
	}

	return user.Email, nil
}
