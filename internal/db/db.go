package db

import (
	"context"
	"fmt"
	"log"

	"github.com/XiaoleC05/CS2Lab/internal/config"
	"github.com/jackc/pgx/v5"
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

	// Test connection
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

// RunMigrations executes SQL migration files
func RunMigrations(migrationFile string) error {
	sql, err := readMigrationFile(migrationFile)
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	_, err = pool.Exec(context.Background(), sql)
	if err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	log.Printf("Migration %s executed successfully", migrationFile)
	return nil
}

func readMigrationFile(filename string) (string, error) {
	// In production, you might want to use embed or read from filesystem
	// For now, we'll use a simple approach
	sql := `
-- Create cs2lab schema
CREATE SCHEMA IF NOT EXISTS cs2lab;

-- Maps table
CREATE TABLE IF NOT EXISTS cs2lab.maps (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    display_name TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- Lineups table
CREATE TABLE IF NOT EXISTS cs2lab.lineups (
    id BIGSERIAL PRIMARY KEY,
    map_id BIGINT NOT NULL REFERENCES cs2lab.maps(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    type TEXT NOT NULL CHECK (type IN ('smoke', 'flash', 'molotov', 'grenade')),
    description TEXT DEFAULT '',
    throw_style TEXT DEFAULT '',
    image_urls TEXT[] DEFAULT '{}',
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Favorites table
CREATE TABLE IF NOT EXISTS cs2lab.favorites (
    user_id BIGINT NOT NULL,
    lineup_id BIGINT NOT NULL REFERENCES cs2lab.lineups(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (user_id, lineup_id)
);

-- Notes table
CREATE TABLE IF NOT EXISTS cs2lab.notes (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    lineup_id BIGINT NOT NULL REFERENCES cs2lab.lineups(id) ON DELETE CASCADE,
    content TEXT DEFAULT '',
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE(user_id, lineup_id)
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_lineups_map_id ON cs2lab.lineups(map_id);
CREATE INDEX IF NOT EXISTS idx_lineups_type ON cs2lab.lineups(type);
CREATE INDEX IF NOT EXISTS idx_favorites_user_id ON cs2lab.favorites(user_id);
CREATE INDEX IF NOT EXISTS idx_notes_user_id ON cs2lab.notes(user_id);
CREATE INDEX IF NOT EXISTS idx_notes_lineup_id ON cs2lab.notes(lineup_id);
`
	return sql, nil
}

// Helper function to handle null arrays
func scanStringArray(row pgx.Row, dest *[]string) error {
	var arr []string
	err := row.Scan(&arr)
	if err != nil {
		return err
	}
	*dest = arr
	return nil
}
