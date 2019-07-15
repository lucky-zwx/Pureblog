# Pureblog  
基于Purecss以及beego打造的博客系统  
**更加快速展示你的内容，只需提前搭建Mysql，无需其它多余操作**  
![ZolGUP.png](https://s2.ax1x.com/2019/07/15/ZolGUP.png)
## 在此之前请在mysql中运行该文件，进行数据库结构导入**导入完成后在admin表中添加用户信息，密码请使用MD5加密后的！**
```
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin`  (
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `pappwd` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `headimg` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`username`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `author` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `top` tinyint(1) NOT NULL DEFAULT 0,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `content` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `morecontent` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `category` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `addtime` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 230 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for session
-- ----------------------------
DROP TABLE IF EXISTS `session`;
CREATE TABLE `session`  (
  `session_key` char(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `session_data` blob NULL,
  `session_expiry` int(11) UNSIGNED NOT NULL,
  PRIMARY KEY (`session_key`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

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
