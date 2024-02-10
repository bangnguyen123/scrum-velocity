package config

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

func connectDB(client *db.PrismaClient) error {
	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	// Disconnect DB before shutting down the application
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		if err := client.Prisma.Disconnect(); err != nil {
			panic(fmt.Errorf("could not disconnect: %w", err))
		}
		os.Exit(0)
	}()
	return nil
}

func InitDB() DB {
	client := db.NewClient()
	if err := connectDB(client); err != nil {
		panic(err)
	}

	ctx := context.Background()

	return DB{
		PrismaClient: client,
		Ctx:          ctx,
	}
}
