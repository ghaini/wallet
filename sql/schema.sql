CREATE TABLE `wallet_changes`
(
    `id`      BIGINT NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL,
    `amount`  DECIMAL(12, 2),
    `created_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB;