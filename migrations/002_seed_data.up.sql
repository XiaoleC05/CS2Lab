-- Seed data for CS2Lab
-- This migration inserts sample CS2 maps and lineups

-- Insert maps
INSERT INTO cs2lab.maps (name, display_name) VALUES
('de_dust2', 'Dust II'),
('de_mirage', 'Mirage'),
('de_inferno', 'Inferno'),
('de_nuke', 'Nuke'),
('de_overpass', 'Overpass'),
('de_ancient', 'Ancient'),
('de_anubis', 'Anubis')
ON CONFLICT (name) DO NOTHING;

-- Insert sample lineups for Dust II
INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT 
    m.id,
    'A Long Smoke',
    'smoke',
    'Smoke for CT spawn from Long A',
    'Left click throw from corner',
    ARRAY['/images/dust2/long_smoke_1.jpg', '/images/dust2/long_smoke_2.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_dust2'
LIMIT 1;

INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT 
    m.id,
    'B Window Flash',
    'flash',
    'Flash pop over B window',
    'Right click jump throw',
    ARRAY['/images/dust2/b_window_flash_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_dust2'
LIMIT 1;

INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT 
    m.id,
    'Mid Doors Molotov',
    'molotov',
    'Molotov for mid doors',
    'Left click throw from T spawn',
    ARRAY['/images/dust2/mid_molly_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_dust2'
LIMIT 1;

-- Insert sample lineups for Mirage
INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT 
    m.id,
    'A Site Smoke from Jungle',
    'smoke',
    'Smoke default A site from jungle',
    'Left click throw',
    ARRAY['/images/mirage/a_smoke_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_mirage'
LIMIT 1;

INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT 
    m.id,
    'Window Flash',
    'flash',
    'Flash for window from T spawn',
    'Jump throw',
    ARRAY['/images/mirage/window_flash_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_mirage'
LIMIT 1;
