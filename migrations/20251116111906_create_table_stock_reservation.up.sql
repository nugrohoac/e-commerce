CREATE TABLE `stock_reservation` (
     `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
     `order_id` BIGINT UNSIGNED,
     `product_id` BIGINT UNSIGNED NOT NULL,
     `warehouse_id` BIGINT UNSIGNED NOT NULL,
     `qty` INT NOT NULL,
     `expires_at` DATETIME NOT NULL,
     `status` ENUM('active', 'released', 'converted_to_order') DEFAULT 'active',
     `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);