-- 建库
CREATE
DATABASE IF NOT EXISTS `user` default charset utf8 COLLATE utf8_general_ci;
-- 切换数据库
use
`user`;

-- 用户表
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `user_id`  INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `phone`    VARCHAR(20)  NOT NULL UNIQUE COMMENT '手机号',
    `name`     VARCHAR(20)  NOT NULL COMMENT '用户名',
    `password` VARCHAR(100) NOT NULL COMMENT 'MD5密码'
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- Student表
DROP TABLE IF EXISTS `students`;
CREATE TABLE `students`
(
    `stu_id`     INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `stu_no`     VARCHAR(20) NOT NULL UNIQUE COMMENT '学号',
    `stu_name`   VARCHAR(50) NOT NULL COMMENT '姓名',
    `stu_gender` VARCHAR(4)  DEFAULT NULL COMMENT '性别',
    `stu_email`  VARCHAR(50) DEFAULT NULL COMMENT '邮箱'
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- 认证码表
DROP TABLE IF EXISTS `auth_codes`;
CREATE TABLE `auth_codes`
(
    `user_id` INT UNSIGNED,
    `stu_id`  INT UNSIGNED,
    `code`    INT UNSIGNED COMMENT '认证码',
    PRIMARY KEY (`user_id`, `stu_id`),
    FOREIGN KEY (`user_id`) REFERENCES users (`user_id`),
    FOREIGN KEY (`stu_id`) REFERENCES students (`stu_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- Session表
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions`
(
    `session_id` VARCHAR(100) PRIMARY KEY,
    `user_id`    INT UNSIGNED NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES users (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;


-- -- 建库
-- CREATE
-- DATABASE IF NOT EXISTS `dorm` default charset utf8 COLLATE utf8_general_ci;
-- -- 切换数据库
-- use
-- `dorm`;

DROP TABLE IF EXISTS `buildings`;
CREATE TABLE `buildings`
(
    `building_id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name`        VARCHAR(50) NOT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

DROP TABLE IF EXISTS `units`;
CREATE TABLE `units`
(
    `unit_id`     INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name`        VARCHAR(50) NOT NULL,
    `building_id` INT UNSIGNED NOT NULL,
    FOREIGN KEY (`building_id`) REFERENCES buildings (`building_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

DROP TABLE IF EXISTS `dorms`;
CREATE TABLE `dorms`
(
    `dorm_id`        INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name`           VARCHAR(20) NOT NULL,
    `gender`         CHAR(10),
    `total_beds`     INT         NOT NULL,
    `available_beds` INT         NOT NULL,
    `unit_id`        INT UNSIGNED NOT NULL,
    FOREIGN KEY (`unit_id`) REFERENCES units (`unit_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;


DROP TABLE IF EXISTS `stu_dorm`;
CREATE TABLE `stu_dorm`
(
    `id`      INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `stu_id`  INT UNSIGNED,
    `dorm_id` INT UNSIGNED,
    FOREIGN KEY (`stu_id`) REFERENCES students (`stu_id`),
    FOREIGN KEY (`dorm_id`) REFERENCES dorms (`dorm_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;


DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`
(
    `order_id`    VARCHAR(100) PRIMARY KEY,
    `user_id`     INT UNSIGNED,
    `count`       INT,
    `building_id` INT UNSIGNED NOT NULL,
    `gender`      VARCHAR(4),
    `create_time` VARCHAR(100),
    `state`       INT UNSIGNED,
    FOREIGN KEY (`user_id`) REFERENCES users (`user_id`),
    FOREIGN KEY (`building_id`) REFERENCES buildings (`building_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`
(
    `id`       INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `order_id` VARCHAR(100),
    `stu_id`   INT UNSIGNED,
    FOREIGN KEY (`stu_id`) REFERENCES students (`stu_id`),
    FOREIGN KEY (`order_id`) REFERENCES orders (`order_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;