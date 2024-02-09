package repositories

import (
	"backend/db"
	"context"
)

type UserRepositoryQ interface {
	GetUsers(user *db.UserModel) []db.UserModel
}

type UserRepository struct {
	ctx    context.Context
	client *db.PrismaClient
}

func NewUserRepository(ctx context.Context, client *db.PrismaClient) *UserRepository {
	return &UserRepository{
		ctx:    ctx,
		client: client,
	}
}

func (userRepository *UserRepository) GetUsers() []db.UserModel {
	users, _ := userRepository.client.User.FindMany().Exec(userRepository.ctx)
	return users
}
