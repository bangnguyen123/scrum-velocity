package user

import (
	"backend/db"
	"backend/repositories"
	"context"
)

type UserServiceWrapper interface {
	GetUsers() error
}

type Service struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(ctx context.Context, prismaClient *db.PrismaClient) *Service {
	userRepository := repositories.NewUserRepository(ctx, prismaClient)
	return &Service{
		UserRepository: userRepository,
	}
}
