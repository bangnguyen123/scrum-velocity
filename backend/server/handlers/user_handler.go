package handlers

import (
	"backend/dtos/responses"
	s "backend/server"
	userService "backend/services/user"

	"github.com/labstack/echo/v4"

	"net/http"
)

type UserHandler struct {
	server *s.Server
}

type User struct {
	ID   string
	Name string
}

func NewUserHandler(server *s.Server) *UserHandler {
	return &UserHandler{server: server}
}

func (userHandler *UserHandler) GetUsers(c echo.Context) error {
	userService := userService.NewUserService(userHandler.server.DB.Ctx, userHandler.server.DB.PrismaClient)
	users := userService.GetUsers()
	return responses.Response(c, http.StatusOK, users)
}
