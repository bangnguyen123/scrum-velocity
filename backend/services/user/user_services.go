package user

import (
	"backend/db"
)

func (userService *Service) GetUsers() []db.UserModel {
	users := userService.UserRepository.GetUsers()
	return users
}
