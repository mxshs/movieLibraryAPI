package services

import (
	"mxshs/movieLibrary/src/domain"
	"mxshs/movieLibrary/src/ports/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (us *UserService) CreateUser(username, password string) (*domain.User, error) {
	return us.userRepo.CreateUser(username, password, domain.USR)
}

func (us *UserService) GetUser(username string) (*domain.User, error) {
	return us.userRepo.GetUserByUsername(username)
}

func (us *UserService) UpdateUser(username, newPassword string) (*domain.User, error) {
	return us.userRepo.UpdateUser(username, newPassword)
}

func (us *UserService) DeleteUser(username string) error {
	return us.userRepo.DeleteUser(username)
}

func (us *UserService) LoginUser(username, password string) (*domain.User, error) {
	return us.userRepo.LoginUser(username, password)
}
