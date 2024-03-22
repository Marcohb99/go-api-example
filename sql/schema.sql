/** TODO add schema */
CREATE DATABASE IF NOT EXISTS `mhb` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `mhb`;

CREATE TABLE IF NOT EXISTS `releases` (
  `uuid` varchar(255) NOT NULL,
  `title` varchar(255) NOT NULL,
  `released` date NULL,
  `resource_url` varchar(255) NULL,
  `uri` varchar(255) NULL,
  `year` varchar(255) NULL,
  PRIMARY KEY (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;