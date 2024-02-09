package config

import (
	"backend/db"
	"context"
)

type DB struct {
	PrismaClient *db.PrismaClient
	Ctx          context.Context
}

func connectDB(client *db.PrismaClient) error {
	if err := client.Prisma.Connect(); err != nil {
		return nil
	}

	// defer func() {
	// 	if err := client.Prisma.Disconnect(); err != nil {
	// 		panic(err)
	// 	}
	// }()
	return nil
}

func InitDB() DB {
	client := db.NewClient()
	connectDB(client)

	ctx := context.Background()

	return DB{
		PrismaClient: client,
		Ctx:          ctx,
	}
}
