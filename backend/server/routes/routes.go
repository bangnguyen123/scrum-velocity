package routes

import (
	s "backend/server"

	"github.com/labstack/echo/v4/middleware"
)

func ConfigRoutes(server *s.Server) {
	ConfigUserRoutes(server)
	server.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
}
