package configs

import (
	"backend/db"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type DB struct {
	PrismaClient *db.PrismaClient
	Ctx          context.Context
}

func ConnectDB(client *db.PrismaClient) error {
	if err := client.Connect(); err != nil {
		return err
	}
	return nil
}

func DisconnectDB(c chan os.Signal, client *db.PrismaClient) {
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		if err := client.Disconnect(); err != nil {
			panic(fmt.Errorf("could not disconnect: %w", err))
		}
		// os.Exit(0)
	}()
}

func InitDB() DB {
	client := db.NewClient()

	if err := ConnectDB(client); err != nil {
		panic(err.Error())
	}

	c := make(chan os.Signal, 1)
	// Disconnect DB before shutting down the application
	DisconnectDB(c, client)

	ctx := context.Background()

	return DB{
		PrismaClient: client,
		Ctx:          ctx,
	}
}
