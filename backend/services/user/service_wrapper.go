package user

type UserServiceWrapper interface {
	GetUsers() error
}

type Service struct {
}

func NewUserService() *Service {
	return &Service{}
}
