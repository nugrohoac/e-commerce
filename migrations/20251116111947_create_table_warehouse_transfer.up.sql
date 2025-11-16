CREATE TABLE `warehouse_transfer` (
  `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `from_warehouse_id` BIGINT UNSIGNED NOT NULL,
  `to_warehouse_id` BIGINT UNSIGNED NOT NULL,
  `product_id` BIGINT UNSIGNED NOT NULL,
  `qty` INT NOT NULL,
  `status` ENUM('pending', 'completed') DEFAULT 'pending',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);