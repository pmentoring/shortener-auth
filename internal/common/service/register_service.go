package service

import (
	"golang.org/x/crypto/bcrypt"
	"shortener-auth/internal/common/repository"
)

type RegisterService struct {
	repo repository.UserRepository
}

func NewRegisterService(repo repository.UserRepository) *RegisterService {
	return &RegisterService{
		repo: repo,
	}
}

func (r *RegisterService) Register(login, password string) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	err = r.repo.CreateUser(login, string(hashedPass))
	if err != nil {
		return err
	}

	return nil
}
