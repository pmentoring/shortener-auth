package service

import (
	"golang.org/x/crypto/bcrypt"
	"shortener-auth/internal/common/repository"
)

type LoginService struct {
	repo       repository.UserRepository
	jwtService *JWTService
}

func (l *LoginService) Login(login, password string) (string, error) {
	user, err := l.repo.GetUserByLogin(login)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	return l.jwtService.CreateJwtForId(user.Id)
}

func NewLoginService(
	repo repository.UserRepository,
	jwtService *JWTService,
) *LoginService {
	return &LoginService{
		repo:       repo,
		jwtService: jwtService,
	}
}
