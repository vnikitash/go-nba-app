# ************************************************************
# Sequel Pro SQL dump
# Version 5446
#
# https://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.32)
# Database: nba
# Generation Time: 2021-06-06 20:56:16 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table players
# ------------------------------------------------------------

DROP TABLE IF EXISTS `players`;

CREATE TABLE `players` (
  `id` varchar(36) NOT NULL DEFAULT '',
  `name` varchar(64) DEFAULT NULL,
  `team_id` varchar(36) DEFAULT NULL,
  `accuracy` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `team` (`team_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table plays
# ------------------------------------------------------------

DROP TABLE IF EXISTS `plays`;

CREATE TABLE `plays` (
  `id` varchar(36) CHARACTER SET utf8 NOT NULL DEFAULT '',
  `team1_id` varchar(36) CHARACTER SET utf8 NOT NULL DEFAULT '',
  `team2_id` varchar(36) CHARACTER SET utf8 NOT NULL DEFAULT '',
  `team1_score` smallint(11) unsigned NOT NULL DEFAULT '0',
  `team2_score` smallint(11) unsigned NOT NULL DEFAULT '0',
  `team1_ast` smallint(11) unsigned NOT NULL DEFAULT '0',
  `team2_ast` smallint(11) unsigned NOT NULL DEFAULT '0',
  `team1_scored2` smallint(11) unsigned DEFAULT '0',
  `team1_scored2_att` smallint(11) unsigned DEFAULT '0',
  `team1_scored3` smallint(11) unsigned DEFAULT '0',
  `team1_scored3_att` smallint(11) unsigned DEFAULT '0',
  `team2_scored2` smallint(11) unsigned DEFAULT '0',
  `team2_scored2_att` smallint(11) unsigned DEFAULT '0',
  `team2_scored3` smallint(11) unsigned DEFAULT '0',
  `team2_scored3_att` smallint(11) unsigned DEFAULT '0',
  KEY `team1_id` (`team1_id`),
  KEY `team2_id` (`team2_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;



# Dump of table teams
# ------------------------------------------------------------

DROP TABLE IF EXISTS `teams`;

CREATE TABLE `teams` (
  `id` varchar(36) NOT NULL DEFAULT '',
  `name` varchar(128) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
