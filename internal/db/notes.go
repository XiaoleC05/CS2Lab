package db

import (
	"context"
	"fmt"

	"github.com/XiaoleC05/CS2Lab/internal/model"
)

// NoteRepository handles note-related database operations
type NoteRepository struct{}

// NewNoteRepository creates a new NoteRepository
func NewNoteRepository() *NoteRepository {
	return &NoteRepository{}
}

// GetByLineup retrieves a user's note for a specific lineup
func (r *NoteRepository) GetByLineup(ctx context.Context, userID, lineupID int64) (*model.Note, error) {
	query := `
		SELECT id, user_id, lineup_id, content, created_at, updated_at
		FROM cs2lab.notes
		WHERE user_id = $1 AND lineup_id = $2
	`

	var n model.Note
	err := pool.QueryRow(ctx, query, userID, lineupID).Scan(&n.ID, &n.UserID, &n.LineupID, &n.Content, &n.CreatedAt, &n.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get note: %w", err)
	}

	return &n, nil
}

// Upsert creates or updates a note
func (r *NoteRepository) Upsert(ctx context.Context, userID, lineupID int64, content string) (*model.Note, error) {
	query := `
		INSERT INTO cs2lab.notes (user_id, lineup_id, content)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, lineup_id)
		DO UPDATE SET content = $3, updated_at = NOW()
		RETURNING id, user_id, lineup_id, content, created_at, updated_at
	`

	var n model.Note
	err := pool.QueryRow(ctx, query, userID, lineupID, content).Scan(&n.ID, &n.UserID, &n.LineupID, &n.Content, &n.CreatedAt, &n.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to upsert note: %w", err)
	}

	return &n, nil
}

// Delete deletes a note and returns affected row count.
func (r *NoteRepository) Delete(ctx context.Context, userID, lineupID int64) (int64, error) {
	query := `
		DELETE FROM cs2lab.notes
		WHERE user_id = $1 AND lineup_id = $2
	`

	result, err := pool.Exec(ctx, query, userID, lineupID)
	if err != nil {
		return 0, fmt.Errorf("failed to delete note: %w", err)
	}

	return result.RowsAffected(), nil
}
