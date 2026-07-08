package db

import (
	"context"
	"fmt"

	"github.com/XiaoleC05/CS2Lab/internal/model"
)

// MapRepository handles map-related database operations
type MapRepository struct{}

// NewMapRepository creates a new MapRepository
func NewMapRepository() *MapRepository {
	return &MapRepository{}
}

// GetAll retrieves all maps with lineup counts
func (r *MapRepository) GetAll(ctx context.Context) ([]model.MapWithLineupCount, error) {
	query := `
		SELECT m.id, m.name, m.display_name, m.created_at, COUNT(l.id) as lineup_count
		FROM cs2lab.maps m
		LEFT JOIN cs2lab.lineups l ON m.id = l.map_id
		GROUP BY m.id, m.name, m.display_name, m.created_at
		ORDER BY m.name
	`

	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query maps: %w", err)
	}
	defer rows.Close()

	var maps []model.MapWithLineupCount
	for rows.Next() {
		var m model.MapWithLineupCount
		err := rows.Scan(&m.ID, &m.Name, &m.DisplayName, &m.CreatedAt, &m.LineupCount)
		if err != nil {
			return nil, fmt.Errorf("failed to scan map: %w", err)
		}
		maps = append(maps, m)
	}

	return maps, rows.Err()
}

// GetByID retrieves a map by ID with lineup count
func (r *MapRepository) GetByID(ctx context.Context, id int64) (*model.MapWithLineupCount, error) {
	query := `
		SELECT m.id, m.name, m.display_name, m.created_at, COUNT(l.id) as lineup_count
		FROM cs2lab.maps m
		LEFT JOIN cs2lab.lineups l ON m.id = l.map_id
		WHERE m.id = $1
		GROUP BY m.id, m.name, m.display_name, m.created_at
	`

	var m model.MapWithLineupCount
	err := pool.QueryRow(ctx, query, id).Scan(&m.ID, &m.Name, &m.DisplayName, &m.CreatedAt, &m.LineupCount)
	if err != nil {
		return nil, fmt.Errorf("failed to get map: %w", err)
	}

	return &m, nil
}
