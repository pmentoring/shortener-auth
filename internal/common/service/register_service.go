package service

import (
	"golang.org/x/crypto/bcrypt"
	"shortener-auth/internal/common/repository"
)

type RegisterService struct {
	repo repository.RegisterRepository
}

func NewRegisterService(repo repository.RegisterRepository) *RegisterService {
	return &RegisterService{
		repo: repo,
	}
}

func (r *RegisterService) Register(login, password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	err = r.repo.CreateUser(login, string(hashedPass))
	if err != nil {
		return "", err
	}

	return "trololo", nil
	// сгенерить жвт
	// вернуть жвт
}