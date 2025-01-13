-- db/migrations/000001_seed_data.down.sql

DELETE FROM foods WHERE restaurant_id IN (1, 2, 3);
DELETE FROM restaurants WHERE id IN (1, 2, 3);