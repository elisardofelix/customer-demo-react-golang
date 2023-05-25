DROP TABLE IF EXISTS `customers`;

CREATE TABLE IF NOT EXISTS `customers`
(
    `id`                  BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `first_name`          VARCHAR(255) NOT NULL,
    `last_name`           VARCHAR(255) NOT NULL,
    `email`               VARCHAR(255) NULL,
    `gender`              VARCHAR(255) NULL,
    `company`             VARCHAR(255) NULL,
    `city`                VARCHAR(255) NULL,
    `title`               VARCHAR(255) NULL,
    `created_at`          TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at`          TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = latin1;

INSERT INTO `customers`(`first_name`,`last_name`,`email`,`gender`,`company`,`city`,`title`)
VALUES ('Elisardo','Felix','elisardo@gmail.com','Male','Verizon', 'Santo Domingo','Sr'),
       ('Juan','Marichal','jma@gmail.com','Male', 'Argenta S.A.','Santiago','Mr');