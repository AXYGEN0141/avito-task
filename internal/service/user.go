package service

import (
	"avito-task/internal"
	"avito-task/internal/model"
)

type UserService struct {
	userRepo internal.UserRepoInterface
}

func NewUserService(repo internal.UserRepoInterface) internal.UserServiceInterface {
	return &UserService{userRepo: repo}
}

// IsExistUser checks if user already exists.
func (s *UserService) IsExistUser(id int) (string, error) {
	return s.userRepo.IsExistUser(id)
}

// CreateUser creates user.
func (a *UserService) CreateUser(user *model.User) (int64, error) {
	return a.userRepo.CreateUser(user)
}
