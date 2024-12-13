package services

import (
	"app/internal/models"
	"app/pkg/bycrypt"
	"app/pkg/middleware"
	"fmt"
)

func (s *ServiceGlobal) RegisterUser(user *models.User) error {
	hashedPassword, err := bycrypt.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	if err := s.CreateUser.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (s *ServiceGlobal) AuthenticateUser(username, password string) (*models.User, error) {
	existingUser, err := s.GetUserByUsername.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if !bycrypt.ComparePassword(existingUser.Password, password) {
		return nil, fmt.Errorf("invalid credentials")
	}

	return existingUser, nil
}

func GenerateAuthToken(user *models.User) (string, error) {
	return middleware.GenerateJWT(user.Id, user.Name, user.Username, user.Password)
}

func (s *ServiceGlobal) GetUserFromToken(token string) (*models.User, error) {
	claims, err := middleware.ParseJWT(token)
	if err != nil {
		return nil, err
	}

	userID, ok := claims["Id"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	user, err := s.GetUserByID.GetUserByID(int(userID))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *ServiceGlobal) FindAllUsers() ([]*models.User, error) {
	return s.GetAllUsers.GetAllUsers()
}
