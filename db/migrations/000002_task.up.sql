BEGIN;

CREATE TABLE `tasks` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `startAt` DATETIME,
  `endAt` DATETIME,
  `status` tinyint NOT NULL,
  `creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP
) ENGINE='MyISAM' COLLATE 'utf8_unicode_ci';

INSERT INTO `tasks` (`name`, `startAt`, `endAt`, `status`)
VALUES ('reading book', '2021-06-30 16:00:49', '2021-06-30 16:00:49', '1');

COMMIT;
