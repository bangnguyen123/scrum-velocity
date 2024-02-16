package main

import (
	// app "backend"
	config "backend/configs"
	"backend/server"
	"backend/server/routes"
	"log"
)

func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	routes.ConfigRoutes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port is already in use")
	}
}

func main() {
	cfg := config.NewConfig()

	Start(cfg)
}
