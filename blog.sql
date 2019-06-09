/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for beego
CREATE DATABASE IF NOT EXISTS `beego` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `beego`;

-- Dumping structure for table beego.admin
CREATE TABLE IF NOT EXISTS `admin` (
  `username` varchar(255) NOT NULL,
  `pappwd` varchar(255) NOT NULL DEFAULT '',
  `headimg` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Dumping data for table beego.admin: ~0 rows (approximately)
/*!40000 ALTER TABLE `admin` DISABLE KEYS */;
INSERT INTO `admin` (`username`, `pappwd`, `headimg`, `email`) VALUES
	('admin', 'e10adc3949ba59abbe56e057f20f883e', '/img/null', 'zhuwx1998@qq.com');
/*!40000 ALTER TABLE `admin` ENABLE KEYS */;

-- Dumping structure for table beego.blog_article
CREATE TABLE IF NOT EXISTS `blog_article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `author` varchar(255) NOT NULL DEFAULT '',
  `top` tinyint(1) NOT NULL DEFAULT '0',
  `title` varchar(255) NOT NULL DEFAULT '',
  `content` longtext NOT NULL,
  `morecontent` longtext NOT NULL,
  `category` varchar(255) NOT NULL DEFAULT '',
  `addtime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table beego.blog_article: ~4 rows (approximately)
/*!40000 ALTER TABLE `blog_article` DISABLE KEYS */;
/*!40000 ALTER TABLE `blog_article` ENABLE KEYS */;

-- Dumping structure for table beego.session
CREATE TABLE IF NOT EXISTS `session` (
  `session_key` char(64) NOT NULL,
  `session_data` blob,
  `session_expiry` int(11) unsigned NOT NULL,
  PRIMARY KEY (`session_key`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Dumping data for table beego.session: 1 rows
/*!40000 ALTER TABLE `session` DISABLE KEYS */;
/*!40000 ALTER TABLE `session` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
