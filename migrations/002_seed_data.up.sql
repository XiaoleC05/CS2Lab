-- Seed data for CS2Lab
-- This migration inserts sample CS2 maps and lineups

-- Insert maps
INSERT INTO cs2lab.maps (name, display_name) VALUES
('de_dust2', '炙热沙城 II'),
('de_mirage', '荒漠迷城'),
('de_inferno', '炼狱小镇'),
('de_nuke', '核子危机'),
('de_overpass', '死亡游乐园'),
('de_ancient', '远古遗迹'),
('de_anubis', '阿努比斯'),
('de_cache', '死城之谜'),
('de_train', '列车停放站')
ON CONFLICT (name) DO NOTHING;

-- Insert sample lineups for Dust II (炙热沙城 II)
INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    'A 长廊烟雾弹',
    'smoke',
    '从 A 长廊封 CT 出生点烟雾，阻断 CT 回防路线',
    '贴左侧墙角，左键投掷',
    ARRAY['/images/dust2/long_smoke_1.jpg', '/images/dust2/long_smoke_2.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_dust2'
LIMIT 1;

INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    'B 窗口闪光弹',
    'flash',
    '从 B 洞外侧向 B 窗口投掷闪光，白屏窗口守敌',
    '右键跳跃投掷',
    ARRAY['/images/dust2/b_window_flash_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_dust2'
LIMIT 1;

INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    '中路门燃烧弹',
    'molotov',
    '从 T 出生点向中路门投掷燃烧弹，清出门后死角',
    '左键投掷，瞄准门框上沿',
    ARRAY['/images/dust2/mid_molly_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_dust2'
LIMIT 1;

-- Insert sample lineups for Mirage (荒漠迷城)
INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    'A 点丛林烟雾弹',
    'smoke',
    '从丛林位置封 A 点默认包点烟雾，配合队友进攻',
    '左键投掷，准星对准屋檐上方',
    ARRAY['/images/mirage/a_smoke_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_mirage'
LIMIT 1;

INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    'VIP 窗口闪光弹',
    'flash',
    '从 T 出生点向 VIP 窗口投掷闪光，白屏中路的 CT 防守',
    '跳跃投掷，准星瞄准窗口中央',
    ARRAY['/images/mirage/window_flash_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_mirage'
LIMIT 1;

-- Insert sample lineups for Cache (死城之谜)
INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    'A 厅烟雾弹',
    'smoke',
    '从 A 厅外侧封 A 包点叉车烟雾，阻断 CT 从 A 小的视线',
    '左键投掷，准星对准仓库边缘',
    ARRAY['/images/cache/a_smoke_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_cache'
LIMIT 1;

INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    '中路闪光弹',
    'flash',
    '从中路车库向警家投掷闪光，白屏中路回防的 CT',
    '右键低抛，反弹墙壁进入中路',
    ARRAY['/images/cache/mid_flash_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_cache'
LIMIT 1;

INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    'B 点燃烧弹',
    'molotov',
    '从 B 厅向 B 包点死角投掷燃烧弹，清出默认躲藏位',
    '左键投掷，瞄准 B 门框上角',
    ARRAY['/images/cache/b_molly_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_cache'
LIMIT 1;

-- Insert sample lineups for Train (列车停放站)
INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    'A 点高低道烟雾弹',
    'smoke',
    '从 T 出生点封 A 点高低道烟雾，隔绝 CT 从红楼梯方向的防守',
    '左键投掷，准星对准高架横梁',
    ARRAY['/images/train/a_smoke_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_train'
LIMIT 1;

INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    '外场闪光弹',
    'flash',
    '从 T 侧通道向外场投掷闪光，白屏外场前压的 CT',
    '右键跳跃投掷，利用车厢反弹',
    ARRAY['/images/train/outside_flash_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_train'
LIMIT 1;

INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    '绿通入口燃烧弹',
    'molotov',
    '从绿通外侧向入口投掷燃烧弹，封锁 CT 从绿通前压的路线',
    '左键投掷，瞄准绿通入口地面',
    ARRAY['/images/train/green_molly_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_train'
LIMIT 1;
