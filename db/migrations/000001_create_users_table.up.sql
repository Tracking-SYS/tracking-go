BEGIN;

CREATE TABLE `products` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `price` int NOT NULL,
  `type` varchar(255) NOT NULL,
  `status` tinyint NOT NULL,
  `creation_time`     DATETIME DEFAULT CURRENT_TIMESTAMP,
  `modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP
) ENGINE='MyISAM' COLLATE 'utf8_unicode_ci';

INSERT INTO `products` (`name`, `price`, `type`, `status`)
VALUES ('Iphone 7', '7000', 'simple', '1');

INSERT INTO `products` (`name`, `price`, `type`, `status`)
VALUES ('Iphone 8', '8000', 'virtual', '2');

INSERT INTO `products` (`name`, `price`, `type`, `status`)
VALUES ('Iphone 9', '9000', 'configurable', '1');

INSERT INTO `products` (`name`, `price`, `type`, `status`)
VALUES ('Iphone 10', '10000', 'simple', '2');

INSERT INTO `products` (`name`, `price`, `type`, `status`)
VALUES ('Iphone 11', '11000', 'bundle', '1');

INSERT INTO `products` (`name`, `price`, `type`, `status`)
VALUES ('Samsung 7', '7000', 'simple', '1');

INSERT INTO `products` (`name`, `price`, `type`, `status`)
VALUES ('Samsung 8', '8000', 'virtual', '2');

INSERT INTO `products` (`name`, `price`, `type`, `status`)
VALUES ('Samsung 9', '9000', 'configurable', '1');
 
INSERT INTO `products` (`name`, `price`, `type`, `status`)
VALUES ('Samsung 10', '10000', 'simple', '2');
 
INSERT INTO `products` (`name`, `price`, `type`, `status`)
VALUES ('Samsung 11', '11000', 'bundle', '1');

COMMIT;
