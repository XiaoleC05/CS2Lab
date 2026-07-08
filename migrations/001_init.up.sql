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

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_lineups_map_id ON cs2lab.lineups(map_id);
CREATE INDEX IF NOT EXISTS idx_lineups_type ON cs2lab.lineups(type);
CREATE INDEX IF NOT EXISTS idx_favorites_user_id ON cs2lab.favorites(user_id);
CREATE INDEX IF NOT EXISTS idx_notes_user_id ON cs2lab.notes(user_id);
CREATE INDEX IF NOT EXISTS idx_notes_lineup_id ON cs2lab.notes(lineup_id);

-- Create full-text search index on lineups
CREATE INDEX IF NOT EXISTS idx_lineups_search ON cs2lab.lineups USING gin(to_tsvector('english', title || ' ' || COALESCE(description, '') || ' ' || COALESCE(throw_style, '')));
