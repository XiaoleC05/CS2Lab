package db

import (
	"context"
	"fmt"

	"github.com/XiaoleC05/CS2Lab/internal/model"
)

// LineupRepository handles lineup-related database operations
type LineupRepository struct{}

// NewLineupRepository creates a new LineupRepository
func NewLineupRepository() *LineupRepository {
	return &LineupRepository{}
}

// GetFiltered retrieves lineups with filters
func (r *LineupRepository) GetFiltered(ctx context.Context, filter model.LineupFilter) ([]model.Lineup, error) {
	query := `
		SELECT id, map_id, title, type, description, throw_style, image_urls, created_at, updated_at
		FROM cs2lab.lineups
		WHERE 1=1
	`
	args := []interface{}{}
	argIndex := 1

	if filter.MapID != nil {
		query += fmt.Sprintf(" AND map_id = $%d", argIndex)
		args = append(args, *filter.MapID)
		argIndex++
	}

	if filter.Type != nil {
		query += fmt.Sprintf(" AND type = $%d", argIndex)
		args = append(args, *filter.Type)
		argIndex++
	}

	if filter.Query != nil && *filter.Query != "" {
		query += fmt.Sprintf(" AND (title ILIKE $%d OR description ILIKE $%d OR throw_style ILIKE $%d)", argIndex, argIndex, argIndex)
		searchQuery := "%" + *filter.Query + "%"
		args = append(args, searchQuery)
		argIndex++
	}

	query += " ORDER BY created_at DESC"

	if filter.Limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argIndex)
		args = append(args, filter.Limit)
		argIndex++
	}

	if filter.Offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", argIndex)
		args = append(args, filter.Offset)
	}

	rows, err := pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query lineups: %w", err)
	}
	defer rows.Close()

	var lineups []model.Lineup
	for rows.Next() {
		var l model.Lineup
		err := rows.Scan(&l.ID, &l.MapID, &l.Title, &l.Type, &l.Description, &l.ThrowStyle, &l.ImageURLs, &l.CreatedAt, &l.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan lineup: %w", err)
		}
		lineups = append(lineups, l)
	}

	return lineups, rows.Err()
}

// GetByID retrieves a lineup by ID
func (r *LineupRepository) GetByID(ctx context.Context, id int64) (*model.Lineup, error) {
	query := `
		SELECT id, map_id, title, type, description, throw_style, image_urls, created_at, updated_at
		FROM cs2lab.lineups
		WHERE id = $1
	`

	var l model.Lineup
	err := pool.QueryRow(ctx, query, id).Scan(&l.ID, &l.MapID, &l.Title, &l.Type, &l.Description, &l.ThrowStyle, &l.ImageURLs, &l.CreatedAt, &l.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get lineup: %w", err)
	}

	return &l, nil
}


// Create adds a new lineup
func (r *LineupRepository) Create(ctx context.Context, l *model.Lineup) (*model.Lineup, error) {
	query := `
		INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, map_id, title, type, description, throw_style, image_urls, created_at, updated_at
	`
	err := pool.QueryRow(ctx, query,
		l.MapID, l.Title, l.Type, l.Description, l.ThrowStyle, l.ImageURLs,
	).Scan(&l.ID, &l.MapID, &l.Title, &l.Type, &l.Description, &l.ThrowStyle, &l.ImageURLs, &l.CreatedAt, &l.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create lineup: %w", err)
	}
	return l, nil
}