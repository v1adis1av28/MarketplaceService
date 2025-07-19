package auth

import (
	"mp-service/internal/models"
	"mp-service/internal/repository/user"
	"mp-service/internal/utils"
)

type AuthService struct {
	repository *user.UserRepository
}

func (a *AuthService) FindUserByEmail(email string) (*models.User, error) {
	return a.repository.FindUserByEmail(email)
}

func NewAuthService(repo *user.UserRepository) *AuthService {
	return &AuthService{repository: repo}
}

func (a *AuthService) CreateUser(info models.RegistrationUserReq) error {
	hashedPassword, err := utils.GenerateHashPassword(info.Password)
	if err != nil {
		return err
	}
	return a.repository.RegistrateUser(info.Email, hashedPassword)
}
