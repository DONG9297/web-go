-- 建库
CREATE
DATABASE IF NOT EXISTS my_db default charset utf8 COLLATE utf8_general_ci;

-- 切换数据库
use
my_db;

-- 建表
DROP TABLE IF EXISTS `users`;

CREATE TABLE `users`
(
    `id`       int(11) unsigned NOT NULL AUTO_INCREMENT,
    `phone`    varchar(100) DEFAULT NULL COMMENT '手机号',
    `name`     varchar(100) DEFAULT NULL COMMENT '用户名',
    `password` varchar(100) DEFAULT NULL COMMENT 'MD5密码',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 插入数据
INSERT INTO `users` (`id`, `phone`, `name`, `password`)
VALUES (NULL, '12345678924', 'dong', 'e10adc3949ba59abbe56e057f20f883e'),
       (NULL, '12345678925', 'dong', 'e10adc3949ba59abbe56e057f20f883e'),
       (NULL, '12345678927', 'dong', 'e10adc3949ba59abbe56e057f20f883e'),
       (NULL, '12345678987', 'dong', 'f30dc57ff8a4d30252d16ed9b97ce272'),
       (NULL, '12345678988', 'dong', 'e10adc3949ba59abbe56e057f20f883e'),
       (NULL, '12345678956', 'dong', 'e10adc3949ba59abbe56e057f20f883e'),
       (NULL, '12345678975', 'dong', 'e10adc3949ba59abbe56e057f20f883e'),
       (NULL, '12345678931', 'dong', 'e10adc3949ba59abbe56e057f20f883e'),
       (NULL, '12345678999', 'dong', 'e10adc3949ba59abbe56e057f20f883e'),
       (NULL, '12345678900', 'dong', 'e10adc3949ba59abbe56e057f20f883e');


-- 建表
DROP TABLE IF EXISTS `dorms`;

CREATE TABLE `dorms`
(
    `dorm_name`             varchar(10) NOT NULL COMMENT '宿舍名',
    `building_name`         varchar(10) DEFAULT NULL COMMENT '楼号',
    `beds_amount`           int         DEFAULT NULL COMMENT '总床位数',
    `availiable_beds_count` int         DEFAULT NULL COMMENT '剩余床位数',
    PRIMARY KEY (`dorm_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 插入数据
INSERT INTO `dorms` (`dorm_name`, `building_name`, `beds_amount`, `availiable_beds_count`)
VALUES ('5001', '5号楼', '4', '1'),
       ('5002', '5号楼', '4', '0'),
       ('5003', '5号楼', '4', '3'),
       ('e1211', '13号楼', '2', '0'),
       ('e1212', '13号楼', '2', '1'),
       ('e1213', '13号楼', '3', '0'),
       ('e1214', '13号楼', '5', '2'),
       ('e1215', '13号楼', '3', '0');

-- 建表
DROP TABLE IF EXISTS `sessions`;

CREATE TABLE `sessions`
(
    `session_id` VARCHAR(100) PRIMARY KEY,
    `user_id`    INT NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES users (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;