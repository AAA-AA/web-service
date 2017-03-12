-- Create syntax for TABLE 'sec_resources'
CREATE TABLE `sec_resources` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL DEFAULT "" COMMENT '菜单名字',
  `resource_key` varchar(200) NOT NULL DEFAULT "" COMMENT '唯一key',
  `type` tinyint(4) NOT NULL DEFAULT -1 COMMENT '类型，0：一级菜单；1：二级菜单；2：三级菜单',
  `uri` varchar(200) NOT NULL DEFAULT "" COMMENT '请求路径',
  `parent_id` int(11) NOT NULL DEFAULT -1 COMMENT '父节点id',
  `level` varchar(200) NOT NULL DEFAULT "",
  `version` int(11) NOT NULL DEFAULT -1,
  `status` tinyint(4) NOT NULL DEFAULT -1,
  `mark` varchar(200) NOT NULL DEFAULT "",
  `operator` int(11) NOT NULL DEFAULT -1,
  `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- Create syntax for TABLE 'sec_role_resource'
CREATE TABLE `sec_role_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `resource_id` int(11) NOT NULL DEFAULT -1 COMMENT '菜单id',
  `role_id` int(11) NOT NULL DEFAULT -1 COMMENT '角色id',
  `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `version` int(11) NOT NULL DEFAULT -1,
  `status` tinyint(4) NOT NULL DEFAULT -1,
  `mark` varchar(200) NOT NULL DEFAULT "",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- Create syntax for TABLE 'sec_role_user'
CREATE TABLE `sec_role_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT -1 COMMENT '用户id',
  `role_id` int(11) NOT NULL DEFAULT -1 COMMENT '角色id',
  `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `version` int(11) NOT NULL DEFAULT -1,
  `status` tinyint(4) NOT NULL DEFAULT -1,
  `mark` varchar(200) NOT NULL DEFAULT "",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- Create syntax for TABLE 'sec_roles'
CREATE TABLE `sec_roles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL DEFAULT "" COMMENT '角色名',
  `role_key` varchar(200) NOT NULL DEFAULT "" COMMENT '角色key',
  `description` varchar(200) NOT NULL DEFAULT "" COMMENT '描述',
  `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `version` int(11) NOT NULL DEFAULT -1,
  `status` tinyint(4) NOT NULL DEFAULT -1,
  `mark` varchar(200) NOT NULL DEFAULT "",
  `operator` int(11) NOT NULL DEFAULT -1,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- Create syntax for TABLE 'sec_users'
CREATE TABLE `sec_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL DEFAULT "" COMMENT '登录用户名',
  `version` int(11) NOT NULL DEFAULT -1,
  `mark` varchar(200) NOT NULL DEFAULT "",
  `operator` int(11) NOT NULL DEFAULT -1,
  `last_login_date` datetime NOT NULL DEFAULT "2017-01-01 00:00:00"  COMMENT '最后登录时间，精确到秒',
   `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '0为正常，1为删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='后台管理员用户表';

