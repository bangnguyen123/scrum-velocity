package user

import (
	dto "backend/dtos/responses"
)

func (userService *Service) GetUsers(user *dto.User) error {
	user.ID = "ID"
	user.Username = "username"
	return nil
}
