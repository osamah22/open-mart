package models

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/osamah22/open-mart/database/migrations"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupTestDB(t *testing.T) (*pgxpool.Pool, func()) {
	t.Helper()
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image: "postgres:15",
		Env: map[string]string{
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_DB":       "testdb",
		},
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
	}

	container, err := testcontainers.GenericContainer(ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		},
	)
	require.NoError(t, err)

	host, err := container.Host(ctx)
	require.NoError(t, err)

	port, err := container.MappedPort(ctx, "5432")
	require.NoError(t, err)

	dsn := fmt.Sprintf(
		"postgres://postgres:password@%s:%s/testdb?sslmode=disable",
		host, port.Port(),
	)

	// ------------------------------------------------
	// 1️⃣ Run migrations using database/sql (Goose)
	// ------------------------------------------------
	cfg, err := pgxpool.ParseConfig(dsn)
	require.NoError(t, err)

	sqlDB := stdlib.OpenDB(*cfg.ConnConfig)

	goose.SetBaseFS(migrations.Files)
	require.NoError(t, goose.Up(sqlDB, "."))

	require.NoError(t, sqlDB.Close())

	// ------------------------------------------------
	// 2️⃣ Create pgxpool for SQLC / runtime
	// ------------------------------------------------
	pool, err := pgxpool.New(ctx, dsn)
	require.NoError(t, err)

	cleanup := func() {
		pool.Close()
		_ = container.Terminate(ctx)
	}

	return pool, cleanup
}
