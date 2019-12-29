# Pureblog  
基于Purecss以及beego打造的博客系统  
[![ln6cMq.md.png](https://s2.ax1x.com/2019/12/29/ln6cMq.md.png)](https://imgchr.com/i/ln6cMq)
**更加快速展示你的内容，只需提前搭建Mysql，无需其它多余操作**  
![ZolGUP.png](https://s2.ax1x.com/2019/07/15/ZolGUP.png)
## 在此之前请在mysql中运行该文件，进行数据库结构导入**导入完成后在admin表中添加用户信息，密码请使用MD5加密后的！**
```
-- --------------------------------------------------------
-- 主机:                           www.xiaoyuan666.com
-- 服务器版本:                        10.0.38-MariaDB - Source distribution
-- 服务器OS:                        Linux
-- HeidiSQL 版本:                  10.2.0.5599
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- Dumping structure for table pureblog.admin
CREATE TABLE IF NOT EXISTS `admin` (
  `username` varchar(255) NOT NULL,
  `pappwd` varchar(255) NOT NULL DEFAULT '',
  `headimg` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Data exporting was unselected.

-- Dumping structure for table pureblog.blog_article
CREATE TABLE IF NOT EXISTS `blog_article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `author` varchar(255) NOT NULL DEFAULT '',
  `top` tinyint(1) NOT NULL DEFAULT '0',
  `title` varchar(255) NOT NULL DEFAULT '',
  `content` longtext NOT NULL,
  `morecontent` longtext NOT NULL,
  `category` varchar(255) NOT NULL DEFAULT '',
  `addtime` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=238 DEFAULT CHARSET=utf8;

-- Data exporting was unselected.

-- Dumping structure for table pureblog.session
CREATE TABLE IF NOT EXISTS `session` (
  `session_key` char(64) NOT NULL,
  `session_data` blob,
  `session_expiry` int(11) unsigned NOT NULL,
  PRIMARY KEY (`session_key`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- Data exporting was unselected.

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;


```
***
+ Linux一键安装方法（现在只支持Linux）
centos/redhat:  
```
yum -y install wget
wget https://github.com/zwx19981207/Pureblog/releases/download/V1.0/install.sh
sudo chmod +x install.sh
sh install.sh
```
ubuntu/deepin/debian:  
```
apt -y install wget
wget https://github.com/zwx19981207/Pureblog/releases/download/V1.0/install.sh
sudo chmod +x install.sh
sh install.sh
```
***
+ 手动安装方法
首先到Release中找到适合你的系统运行文件，点击下载  
![ZoMDMj.png](https://s2.ax1x.com/2019/07/15/ZoMDMj.png)  
接下来进行解压，在conf文件夹中添加xapp.conf文件  
文件类容如下  
```
blog_name = "请将这里修改为你的博客标题"
blog_second_name = "你的座右铭"
blog_url = "你的网站地址"
navitem_github_link = "你的giehub页面地址"
navitem_githubio_link = "你的github io页面地址"
xsrfkey = "Htr5RPlsU+KDg5mW5DwLHC2RCPBBaRpQdY0s0/loW9uzZk/gUJp4pA=="
sessionproviderconfig = "数据库用户名:数据库密码@tcp(数据库地址记得加端口)/zhuwx?charset=utf8&loc=Local"
sessionsavepath = "数据库用户名:数据库密码@tcp(数据库地址记得加端口)/zhuwx?charset=utf8&loc=Local"
sessionhashkey = "My+XbzFY74b8QUsciYZxrmOvfXyF"
mysqlurls = "数据库用户名:数据库密码@tcp(数据库地址记得加端口)/zhuwx?charset=utf8&loc=Local"
```
接下来到解压后的根目录运行Pureblog文件（记得增加运行权限）  
如需后台运行，请使用以下代码运行：  
```
nohup ./Purelbog &
```
***
+ app.conf配置文件  
httpport为运行端口，可作为nginx反向代理到80端口。如果想直接使用可以改到**80！！**
