CREATE TABLE line_foods (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    food_id BIGINT NOT NULL,
    restaurant_id BIGINT NOT NULL,
    order_id BIGINT DEFAULT NULL,
    count INT NOT NULL DEFAULT 0,
    active BOOLEAN NOT NULL DEFAULT FALSE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_foods_line_food FOREIGN KEY (food_id) REFERENCES foods (id),
    CONSTRAINT fk_restaurants_line_food FOREIGN KEY (restaurant_id) REFERENCES restaurants (id),
    CONSTRAINT fk_orders_line_food FOREIGN KEY (order_id) REFERENCES orders (id)
);