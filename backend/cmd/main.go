package main

import (
	app "backend"
	config "backend/configs"
)

func main() {
	cfg := config.NewConfig()

	app.Start(cfg)
}
