package postgres

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joe07leovin/Bao/internal/platform/config"
)

func TestRejectPlaintextBeforeConnecting(t *testing.T) {
	_, err := Open(context.Background(), config.Config{DatabaseURL: "postgres://bao@localhost/bao?sslmode=disable", DatabaseMaxConnections: 1, DatabaseConnectTimeout: time.Second})
	if err == nil {
		t.Fatal("plaintext connection accepted")
	}
}
func TestIntegrationVerifiedTLS(t *testing.T) {
	dsn := os.Getenv("BAO_TEST_DATABASE_URL")
	if dsn == "" {
		t.Skip("set BAO_TEST_DATABASE_URL to an isolated PostgreSQL URL using sslmode=verify-full")
	}
	pool, err := Open(context.Background(), config.Config{DatabaseURL: dsn, DatabaseMaxConnections: 1, DatabaseConnectTimeout: time.Second})
	if err != nil {
		t.Fatal(err)
	}
	pool.Close()
}
