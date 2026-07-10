-- 003_lineup_unique: Add UNIQUE constraint on lineups(map_id, title)
-- Idempotent: safe to run multiple times on existing databases

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint
        WHERE conname = 'unique_lineup_map_title'
          AND conrelid = 'cs2lab.lineups'::regclass
    ) THEN
        ALTER TABLE cs2lab.lineups
            ADD CONSTRAINT unique_lineup_map_title UNIQUE (map_id, title);
    END IF;
END $$;