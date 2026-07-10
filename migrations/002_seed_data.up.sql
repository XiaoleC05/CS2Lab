-- Seed data for CS2Lab
-- This migration inserts sample CS2 maps and lineups

-- Insert maps
INSERT INTO cs2lab.maps (name, display_name) VALUES
('de_dust2', '鐐欑儹娌欏煄 II'),
('de_mirage', '鑽掓紶杩峰煄'),
('de_inferno', '鐐肩嫳灏忛晣'),
('de_nuke', '鏍稿瓙鍗辨満'),
('de_overpass', '姝讳骸娓镐箰鍥?),
('de_ancient', '杩滃彜閬楄抗'),
('de_anubis', '闃垮姫姣旀柉'),
('de_cache', '姝诲煄涔嬭皽'),
('de_train', '鍒楄溅鍋滄斁绔?)
ON CONFLICT (name) DO UPDATE SET display_name = EXCLUDED.display_name;

-- Insert sample lineups for Dust II (鐐欑儹娌欏煄 II)
INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    'A 闀垮粖鐑熼浘寮?,
    'smoke',
    '浠?A 闀垮粖灏?CT 鍑虹敓鐐圭儫闆撅紝闃绘柇 CT 鍥為槻璺嚎',
    '璐村乏渚у瑙掞紝宸﹂敭鎶曟幏',
    ARRAY['/images/dust2/long_smoke_1.jpg', '/images/dust2/long_smoke_2.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_dust2'
  AND NOT EXISTS (
      SELECT 1 FROM cs2lab.lineups l WHERE l.map_id = m.id AND l.title = 'A 闀垮粖鐑熼浘寮?
  )
LIMIT 1;

INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    'B 绐楀彛闂厜寮?,
    'flash',
    '浠?B 娲炲渚у悜 B 绐楀彛鎶曟幏闂厜锛岀櫧灞忕獥鍙ｅ畧鏁?,
    '鍙抽敭璺宠穬鎶曟幏',
    ARRAY['/images/dust2/b_window_flash_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_dust2'
  AND NOT EXISTS (
      SELECT 1 FROM cs2lab.lineups l WHERE l.map_id = m.id AND l.title = 'B 绐楀彛闂厜寮?
  )
LIMIT 1;

INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    '涓矾闂ㄧ噧鐑у脊',
    'molotov',
    '浠?T 鍑虹敓鐐瑰悜涓矾闂ㄦ姇鎺风噧鐑у脊锛屾竻鍑洪棬鍚庢瑙?,
    '宸﹂敭鎶曟幏锛岀瀯鍑嗛棬妗嗕笂娌?,
    ARRAY['/images/dust2/mid_molly_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_dust2'
  AND NOT EXISTS (
      SELECT 1 FROM cs2lab.lineups l WHERE l.map_id = m.id AND l.title = '涓矾闂ㄧ噧鐑у脊'
  )
LIMIT 1;

-- Insert sample lineups for Mirage (鑽掓紶杩峰煄)
INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    'A 鐐逛笡鏋楃儫闆惧脊',
    'smoke',
    '浠庝笡鏋椾綅缃皝 A 鐐归粯璁ゅ寘鐐圭儫闆撅紝閰嶅悎闃熷弸杩涙敾',
    '宸﹂敭鎶曟幏锛屽噯鏄熷鍑嗗眿妾愪笂鏂?,
    ARRAY['/images/mirage/a_smoke_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_mirage'
  AND NOT EXISTS (
      SELECT 1 FROM cs2lab.lineups l WHERE l.map_id = m.id AND l.title = 'A 鐐逛笡鏋楃儫闆惧脊'
  )
LIMIT 1;

INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    'VIP 绐楀彛闂厜寮?,
    'flash',
    '浠?T 鍑虹敓鐐瑰悜 VIP 绐楀彛鎶曟幏闂厜锛岀櫧灞忎腑璺殑 CT 闃插畧',
    '璺宠穬鎶曟幏锛屽噯鏄熺瀯鍑嗙獥鍙ｄ腑澶?,
    ARRAY['/images/mirage/window_flash_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_mirage'
  AND NOT EXISTS (
      SELECT 1 FROM cs2lab.lineups l WHERE l.map_id = m.id AND l.title = 'VIP 绐楀彛闂厜寮?
  )
LIMIT 1;

-- Insert sample lineups for Cache (姝诲煄涔嬭皽)
INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    'A 鍘呯儫闆惧脊',
    'smoke',
    '浠?A 鍘呭渚у皝 A 鍖呯偣鍙夎溅鐑熼浘锛岄樆鏂?CT 浠?A 灏忕殑瑙嗙嚎',
    '宸﹂敭鎶曟幏锛屽噯鏄熷鍑嗕粨搴撹竟缂?,
    ARRAY['/images/cache/a_smoke_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_cache'
  AND NOT EXISTS (
      SELECT 1 FROM cs2lab.lineups l WHERE l.map_id = m.id AND l.title = 'A 鍘呯儫闆惧脊'
  )
LIMIT 1;

INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    '涓矾闂厜寮?,
    'flash',
    '浠庝腑璺溅搴撳悜璀﹀鎶曟幏闂厜锛岀櫧灞忎腑璺洖闃茬殑 CT',
    '鍙抽敭浣庢姏锛屽弽寮瑰澹佽繘鍏ヤ腑璺?,
    ARRAY['/images/cache/mid_flash_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_cache'
  AND NOT EXISTS (
      SELECT 1 FROM cs2lab.lineups l WHERE l.map_id = m.id AND l.title = '涓矾闂厜寮?
  )
LIMIT 1;

INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    'B 鐐圭噧鐑у脊',
    'molotov',
    '浠?B 鍘呭悜 B 鍖呯偣姝昏鎶曟幏鐕冪儳寮癸紝娓呭嚭榛樿韬茶棌浣?,
    '宸﹂敭鎶曟幏锛岀瀯鍑?B 闂ㄦ涓婅',
    ARRAY['/images/cache/b_molly_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_cache'
  AND NOT EXISTS (
      SELECT 1 FROM cs2lab.lineups l WHERE l.map_id = m.id AND l.title = 'B 鐐圭噧鐑у脊'
  )
LIMIT 1;

-- Insert sample lineups for Train (鍒楄溅鍋滄斁绔?
INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    'A 鐐归珮浣庨亾鐑熼浘寮?,
    'smoke',
    '浠?T 鍑虹敓鐐瑰皝 A 鐐归珮浣庨亾鐑熼浘锛岄殧缁?CT 浠庣孩妤兼鏂瑰悜鐨勯槻瀹?,
    '宸﹂敭鎶曟幏锛屽噯鏄熷鍑嗛珮鏋舵í姊?,
    ARRAY['/images/train/a_smoke_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_train'
  AND NOT EXISTS (
      SELECT 1 FROM cs2lab.lineups l WHERE l.map_id = m.id AND l.title = 'A 鐐归珮浣庨亾鐑熼浘寮?
  )
LIMIT 1;

INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    '澶栧満闂厜寮?,
    'flash',
    '浠?T 渚ч€氶亾鍚戝鍦烘姇鎺烽棯鍏夛紝鐧藉睆澶栧満鍓嶅帇鐨?CT',
    '鍙抽敭璺宠穬鎶曟幏锛屽埄鐢ㄨ溅鍘㈠弽寮?,
    ARRAY['/images/train/outside_flash_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_train'
  AND NOT EXISTS (
      SELECT 1 FROM cs2lab.lineups l WHERE l.map_id = m.id AND l.title = '澶栧満闂厜寮?
  )
LIMIT 1;

INSERT INTO cs2lab.lineups (map_id, title, type, description, throw_style, image_urls)
SELECT
    m.id,
    '缁块€氬叆鍙ｇ噧鐑у脊',
    'molotov',
    '浠庣豢閫氬渚у悜鍏ュ彛鎶曟幏鐕冪儳寮癸紝灏侀攣 CT 浠庣豢閫氬墠鍘嬬殑璺嚎',
    '宸﹂敭鎶曟幏锛岀瀯鍑嗙豢閫氬叆鍙ｅ湴闈?,
    ARRAY['/images/train/green_molly_1.jpg']
FROM cs2lab.maps m
WHERE m.name = 'de_train'
  AND NOT EXISTS (
      SELECT 1 FROM cs2lab.lineups l WHERE l.map_id = m.id AND l.title = '缁块€氬叆鍙ｇ噧鐑у脊'
  )
LIMIT 1;