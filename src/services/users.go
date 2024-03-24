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

func (us *UserService) GetUser(id int) (*domain.User, error) {
	return us.userRepo.GetUser(id)
}

func (us *UserService) GetUsers() ([]*domain.User, error) {
	return us.userRepo.GetUsers()
}

func (us *UserService) UpdateUser(id int, newUsername, newPassword string) (*domain.User, error) {
	return us.userRepo.UpdateUser(id, newUsername, newPassword)
}

func (us *UserService) DeleteUser(id int) error {
	return us.userRepo.DeleteUser(id)
}

func (us *UserService) LoginUser(username, password string) (*domain.User, error) {
	return us.userRepo.LoginUser(username, password)
}
