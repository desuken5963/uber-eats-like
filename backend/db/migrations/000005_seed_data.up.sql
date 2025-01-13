-- db/migrations/000001_seed_data.up.sql

INSERT INTO restaurants (name, fee, time_required, created_at, updated_at)
VALUES 
    ('testレストラン_0', 100, 10, NOW(), NOW()),
    ('testレストラン_1', 100, 10, NOW(), NOW()),
    ('testレストラン_2', 100, 10, NOW(), NOW());

-- レストランIDに紐付けてフードデータを挿入
INSERT INTO foods (restaurant_id, name, price, description, created_at, updated_at)
VALUES 
    (1, 'フード名_0', 500, 'フード_0の説明文です。', NOW(), NOW()),
    (1, 'フード名_1', 500, 'フード_1の説明文です。', NOW(), NOW()),
    (1, 'フード名_2', 500, 'フード_2の説明文です。', NOW(), NOW()),
    (1, 'フード名_3', 500, 'フード_3の説明文です。', NOW(), NOW()),
    (1, 'フード名_4', 500, 'フード_4の説明文です。', NOW(), NOW()),
    (1, 'フード名_5', 500, 'フード_5の説明文です。', NOW(), NOW()),
    (2, 'フード名_0', 500, 'フード_0の説明文です。', NOW(), NOW()),
    (2, 'フード名_1', 500, 'フード_1の説明文です。', NOW(), NOW()),
    (2, 'フード名_2', 500, 'フード_2の説明文です。', NOW(), NOW()),
    (3, 'フード名_0', 500, 'フード_0の説明文です。', NOW(), NOW()),
    (3, 'フード名_1', 500, 'フード_1の説明文です。', NOW(), NOW()),
    (3, 'フード名_2', 500, 'フード_2の説明文です。', NOW(), NOW());