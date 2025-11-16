CREATE TABLE `order` (
     `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
     `user_id` BIGINT UNSIGNED NOT NULL,
     `status` ENUM('pending', 'paid', 'expired', 'canceled') DEFAULT 'pending',
     `total_price` DECIMAL(12,2) NOT NULL DEFAULT 0,
     `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);