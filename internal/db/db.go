package db

import (
	"context"
	"fmt"
	"log"

	"github.com/XiaoleC05/CS2Lab/internal/config"
	"github.com/XiaoleC05/CS2Lab/migrations"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

// Init initializes the database connection pool
func Init(cfg *config.Config) error {
	var err error
	pool, err = pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return fmt.Errorf("unable to ping database: %w", err)
	}

	log.Println("Database connection established")
	return nil
}

// Close closes the database connection pool
func Close() {
	if pool != nil {
		pool.Close()
		log.Println("Database connection closed")
	}
}

// GetPool returns the database connection pool
func GetPool() *pgxpool.Pool {
	return pool
}

func ensureMigrationTable(ctx context.Context) error {
	_, err := pool.Exec(ctx, `
		CREATE SCHEMA IF NOT EXISTS cs2lab;
		CREATE TABLE IF NOT EXISTS cs2lab.schema_migrations (
			version TEXT PRIMARY KEY,
			applied_at TIMESTAMPTZ DEFAULT NOW()
		);
	`)
	return err
}

func isMigrationApplied(ctx context.Context, version string) (bool, error) {
	var applied bool
	err := pool.QueryRow(ctx,
		`SELECT EXISTS(SELECT 1 FROM cs2lab.schema_migrations WHERE version = $1)`,
		version,
	).Scan(&applied)
	return applied, err
}

// RunMigrations executes a SQL migration file from migrations/*.up.sql once.
func RunMigrations(migrationName string) error {
	ctx := context.Background()

	if err := ensureMigrationTable(ctx); err != nil {
		return fmt.Errorf("failed to ensure migration table: %w", err)
	}

	applied, err := isMigrationApplied(ctx, migrationName)
	if err != nil {
		return fmt.Errorf("failed to check migration status: %w", err)
	}
	if applied {
		log.Printf("Migration %s already applied, skipping", migrationName)
		return nil
	}

	filename := migrationName + ".up.sql"
	sqlBytes, err := migrations.Files.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read migration file %s: %w", filename, err)
	}

	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin migration transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, string(sqlBytes)); err != nil {
		return fmt.Errorf("failed to execute migration %s: %w", migrationName, err)
	}

	if _, err := tx.Exec(ctx,
		`INSERT INTO cs2lab.schema_migrations (version) VALUES ($1)`,
		migrationName,
	); err != nil {
		return fmt.Errorf("failed to record migration %s: %w", migrationName, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit migration %s: %w", migrationName, err)
	}

	log.Printf("Migration %s executed successfully", migrationName)
	return nil
}
