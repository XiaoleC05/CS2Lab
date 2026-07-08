package model

import "time"

// Map represents a CS2 map
type Map struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"display_name"`
	CreatedAt   time.Time `json:"created_at"`
}

// Lineup represents a utility lineup (smoke, flash, molotov, grenade)
type Lineup struct {
	ID          int64     `json:"id"`
	MapID       int64     `json:"map_id"`
	Title       string    `json:"title"`
	Type        string    `json:"type"` // smoke, flash, molotov, grenade
	Description string    `json:"description"`
	ThrowStyle  string    `json:"throw_style"`
	ImageURLs   []string  `json:"image_urls"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Favorite represents a user's favorite lineup
type Favorite struct {
	UserID    int64     `json:"user_id"`
	LineupID  int64     `json:"lineup_id"`
	CreatedAt time.Time `json:"created_at"`
}

// Note represents a user's personal note on a lineup
type Note struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	LineupID  int64     `json:"lineup_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MapWithLineupCount represents a map with the count of lineups
type MapWithLineupCount struct {
	Map
	LineupCount int `json:"lineup_count"`
}

// LineupFilter represents filters for lineup queries
type LineupFilter struct {
	MapID  *int64
	Type   *string
	Query  *string
	Limit  int
	Offset int
}
