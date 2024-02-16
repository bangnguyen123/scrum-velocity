package db

import (
	"context"
	"os"
	"testing"
	"time"

	configs "backend/configs"
	prismaDB "backend/db"
	setupTest "backend/test"

	"github.com/steebchen/prisma-client-go/test"
	"github.com/steebchen/prisma-client-go/test/helpers/massert"
	"github.com/stretchr/testify/assert"
)

type cx = context.Context
type Func func(t *testing.T, db test.Database, client *prismaDB.PrismaClient, ctx cx)

func TestDBConnection(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		before []string
		run    Func
	}{{
		name: "connect success",
		run: func(t *testing.T, db test.Database, client *prismaDB.PrismaClient, ctx cx) {
			// manually setup testing
			mockDBName := db.SetupDatabase(t)
			setupTest.Migrate(t, db, client.Engine, mockDBName)

			defer test.Teardown(t, db, mockDBName)

			configs.ConnectDB(client)

			if err := client.Disconnect(); err != nil {
				t.Fatalf("fail %s", err)
			}
		},
	}, {
		name: "not connected yet",
		run: func(t *testing.T, db test.Database, client *prismaDB.PrismaClient, ctx cx) {

			// manually setup testing
			mockDBName := db.SetupDatabase(t)
			setupTest.Migrate(t, db, client.Engine, mockDBName)

			defer test.Teardown(t, db, mockDBName)

			// The order of params must follow the order of defined schema
			_, err := client.User.CreateOne(
				prismaDB.User.Name.Set("name"),
				prismaDB.User.Username.Set("username"),
				prismaDB.User.Password.Set("password"),
			).Exec(ctx)

			assert.NotEqual(t, err, nil)
			massert.Equal(t, "request failed: client is not connected yet", err.Error())
		},
	}, {
		name: "already disconnected",
		run: func(t *testing.T, db test.Database, client *prismaDB.PrismaClient, ctx cx) {
			// manually setup testing
			mockDBName := db.SetupDatabase(t)
			setupTest.Migrate(t, db, client.Engine, mockDBName)

			defer test.Teardown(t, db, mockDBName)

			configs.ConnectDB(client)

			_, err := client.User.CreateOne(
				prismaDB.User.Name.Set("name"),
				prismaDB.User.Username.Set("username"),
				prismaDB.User.Password.Set("password"),
			).Exec(ctx)

			if err != nil {
				t.Fatalf("fail %s", err)
			}

			if err := client.Disconnect(); err != nil {
				t.Fatalf("fail %s", err)
			}

			_, err = client.User.CreateOne(
				prismaDB.User.Name.Set("name"),
				prismaDB.User.Username.Set("username"),
				prismaDB.User.Password.Set("password"),
			).Exec(ctx)

			assert.NotEqual(t, err, nil)
			massert.Equal(t, "request failed: client is already disconnected", err.Error())
		},
	}, {
		name: "connect err on async query engine error",
		run: func(t *testing.T, db test.Database, client *prismaDB.PrismaClient, ctx cx) {
			err := configs.ConnectDB(client)

			assert.Regexp(t, "Environment variable not found: DATABASE_URL", err.Error())
		},
	}, {
		name: "Disconnect DB in Go Channel",
		run: func(t *testing.T, db test.Database, client *prismaDB.PrismaClient, ctx cx) {
			mockDBName := db.SetupDatabase(t)
			setupTest.Migrate(t, db, client.Engine, mockDBName)

			defer test.Teardown(t, db, mockDBName)

			configs.ConnectDB(client)

			signalChan := make(chan os.Signal, 1)
			configs.DisconnectDB(signalChan, client)

			// Simulate an interrupt signal to trigger the shutdown
			interruptSignal := os.Interrupt

			// signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

			signalChan <- interruptSignal

			// Wait for the shutdown to complete
			<-time.After(1000 * time.Millisecond)
			_, err := client.User.CreateOne(
				prismaDB.User.Name.Set("name"),
				prismaDB.User.Username.Set("username"),
				prismaDB.User.Password.Set("password"),
			).Exec(ctx)

			massert.Equal(t, "request failed: client is already disconnected", err.Error())
		},
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			test.RunSerial(t, []test.Database{test.PostgreSQL}, func(t *testing.T, db test.Database, ctx context.Context) {
				client := prismaDB.NewClient()
				tt.run(t, db, client, context.Background())
			})
		})
	}
}
