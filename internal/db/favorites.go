package db

import (
	"context"
	"fmt"

	"github.com/XiaoleC05/CS2Lab/internal/model"
)

// FavoriteRepository handles favorite-related database operations
type FavoriteRepository struct{}

// NewFavoriteRepository creates a new FavoriteRepository
func NewFavoriteRepository() *FavoriteRepository {
	return &FavoriteRepository{}
}

// GetByUser retrieves all favorites for a user
func (r *FavoriteRepository) GetByUser(ctx context.Context, userID int64) ([]model.Favorite, error) {
	query := `
		SELECT user_id, lineup_id, created_at
		FROM cs2lab.favorites
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	rows, err := pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query favorites: %w", err)
	}
	defer rows.Close()

	var favorites []model.Favorite
	for rows.Next() {
		var f model.Favorite
		err := rows.Scan(&f.UserID, &f.LineupID, &f.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan favorite: %w", err)
		}
		favorites = append(favorites, f)
	}

	return favorites, rows.Err()
}

// Add adds a lineup to user's favorites
func (r *FavoriteRepository) Add(ctx context.Context, userID, lineupID int64) error {
	query := `
		INSERT INTO cs2lab.favorites (user_id, lineup_id)
		VALUES ($1, $2)
		ON CONFLICT (user_id, lineup_id) DO NOTHING
	`

	_, err := pool.Exec(ctx, query, userID, lineupID)
	if err != nil {
		return fmt.Errorf("failed to add favorite: %w", err)
	}

	return nil
}

// Remove removes a lineup from user's favorites
func (r *FavoriteRepository) Remove(ctx context.Context, userID, lineupID int64) error {
	query := `
		DELETE FROM cs2lab.favorites
		WHERE user_id = $1 AND lineup_id = $2
	`

	_, err := pool.Exec(ctx, query, userID, lineupID)
	if err != nil {
		return fmt.Errorf("failed to remove favorite: %w", err)
	}

	return nil
}
