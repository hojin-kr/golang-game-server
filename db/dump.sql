grant all privileges on user.* to app@'%' IDENTIFIED BY '1q2w3e4r';
grant all privileges on user.* to app@'localhost' IDENTIFIED BY '1q2w3e4r';
CREATE DATABASE IF NOT EXISTS user default charset utf8;
use user;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `platform_id` varchar(128) DEFAULT NULL,
  `platform` varchar(10) DEFAULT NULL,
  `created` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `device_id` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `platform_id` (`platform_id`,`platform`)
);