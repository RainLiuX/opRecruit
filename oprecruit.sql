SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";
CREATE TABLE `user_info` (
	`id` int(10) NOT NULL PRIMARY KEY AUTO_INCREMENT,
	`netid` char(20) NOT NULL,
	`number` char(11) NOT NULL,
	`nickname` char(20) NOT NULL,
	`pwd` char(32) NOT NULL,
	`grade` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `container_info` (
	`id` int(10) NOT NULL PRIMARY KEY AUTO_INCREMENT,
	`dockerid` char(32) NOT NULL,
	`staus` char(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
