package routes

import (
	"backend/server"
	"backend/server/handlers"
)

func ConfigUserRoutes(server *server.Server) {
	userHandler := handlers.NewUserHandler(server)
	server.Echo.GET("/users", userHandler.GetUsers)
}
