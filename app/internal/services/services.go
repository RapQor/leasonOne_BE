package services

import "app/internal/repositories"

type ServiceGlobal struct {
	CreateUser        repositories.UserRepo
	CheckExistingUser repositories.UserRepo
	GetUserByID       repositories.UserRepo
	GetUserByUsername repositories.UserRepo
	GetAllUsers       repositories.UserRepo
}

func Service(service ServiceGlobal) *ServiceGlobal {
	return &ServiceGlobal{
		CreateUser:        service.CreateUser,
		CheckExistingUser: service.CheckExistingUser,
		GetUserByID:       service.GetUserByID,
		GetUserByUsername: service.GetUserByUsername,
		GetAllUsers:       service.GetAllUsers,
	}
}
