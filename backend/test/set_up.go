package test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/steebchen/prisma-client-go/cli"
	"github.com/steebchen/prisma-client-go/engine"
	"github.com/steebchen/prisma-client-go/test"
	"github.com/steebchen/prisma-client-go/test/cmd"
)

const schemaTemplate = "schema.temp.%s.prisma"

func Migrate(t *testing.T, db test.Database, e engine.Engine, mockDB string) {
	t.Helper()
	schemaPath := fmt.Sprintf(schemaTemplate, db.Name())

	xe := e.(*engine.QueryEngine)
	xe.ReplaceSchema(func(schema string) string {
		for _, fromDB := range test.Databases {
			schema = strings.ReplaceAll(schema, fmt.Sprintf(`"%s"`, fromDB.Name()), fmt.Sprintf(`"%s"`, db.Name()))
		}
		return schema
	})
	xe.ReplaceSchema(func(schema string) string {
		return strings.ReplaceAll(schema, `env("DATABASE_URL")`, fmt.Sprintf(`"%s"`, db.ConnectionString(mockDB)))
	})
	if err := os.WriteFile(schemaPath, []byte(xe.Schema), 0644); err != nil {
		t.Fatal(err)
	}

	runDBPush(t, schemaPath)
}

func runDBPush(t *testing.T, schemaPath string) {
	cleanup(t)

	verbose := os.Getenv("PRISMA_CLIENT_GO_TEST_MIGRATE_LOGS") != ""
	if err := cli.Run([]string{"db", "push", "--schema=./" + schemaPath, "--skip-generate"}, verbose); err != nil {
		t.Fatalf("could not run db push: %s", err)
	}
}

func cleanup(t *testing.T) {
	if err := cmd.Run("rm", "-rf", "migrations"); err != nil {
		t.Fatal(err)
	}

	if err := cmd.Run("rm", "-rf", "*.db"); err != nil {
		t.Fatal(err)
	}
}
