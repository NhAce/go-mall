-- 创建数据库
CREATE DATABASE IF NOT EXISTS go_mall CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE go_mall;

-- 用户表
CREATE TABLE IF NOT EXISTS `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `openid` varchar(64) NOT NULL COMMENT '微信openid',
  `nickname` varchar(64) DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机号',
  `gender` tinyint DEFAULT '0' COMMENT '性别 0:未知 1:男 2:女',
  `country` varchar(64) DEFAULT NULL COMMENT '国家',
  `province` varchar(64) DEFAULT NULL COMMENT '省份',
  `city` varchar(64) DEFAULT NULL COMMENT '城市',
  `language` varchar(16) DEFAULT NULL COMMENT '语言',
  `status` tinyint DEFAULT '1' COMMENT '状态 1:正常 0:禁用',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_openid` (`openid`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 商品分类表
CREATE TABLE IF NOT EXISTS `category` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL COMMENT '分类名称',
  `sort` int DEFAULT '0' COMMENT '排序',
  `status` tinyint DEFAULT '1' COMMENT '状态 1:启用 0:禁用',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品分类表';

-- 商品表
CREATE TABLE IF NOT EXISTS `product` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL COMMENT '商品名称',
  `category_id` int unsigned NOT NULL COMMENT '分类ID',
  `price` decimal(10,2) NOT NULL COMMENT '售价',
  `stock` int NOT NULL DEFAULT '0' COMMENT '库存',
  `cover_img` varchar(255) DEFAULT NULL COMMENT '封面图片',
  `images` text COMMENT '商品图片 JSON格式',
  `description` text COMMENT '商品描述',
  `status` tinyint DEFAULT '1' COMMENT '状态 1:上架 0:下架',
  `sort` int DEFAULT '0' COMMENT '排序',
  `sales` int DEFAULT '0' COMMENT '销量',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品表';

-- 购物车表
CREATE TABLE IF NOT EXISTS `cart` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `product_id` int unsigned NOT NULL COMMENT '商品ID',
  `quantity` int NOT NULL DEFAULT '1' COMMENT '数量',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_product_id` (`product_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='购物车表';

-- 收货地址表
CREATE TABLE IF NOT EXISTS `address` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `name` varchar(64) NOT NULL COMMENT '收货人姓名',
  `phone` varchar(20) NOT NULL COMMENT '手机号',
  `province` varchar(64) NOT NULL COMMENT '省',
  `city` varchar(64) NOT NULL COMMENT '市',
  `district` varchar(64) NOT NULL COMMENT '区',
  `detail` varchar(255) NOT NULL COMMENT '详细地址',
  `is_default` tinyint(1) DEFAULT '0' COMMENT '是否默认地址',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='收货地址表';

-- 订单表
CREATE TABLE IF NOT EXISTS `order` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `order_no` varchar(64) NOT NULL COMMENT '订单号',
  `total_price` decimal(10,2) NOT NULL COMMENT '总价',
  `status` tinyint DEFAULT '0' COMMENT '订单状态 0:待付款 1:待发货 2:待收货 3:已完成 4:已取消',
  `address_id` int unsigned NOT NULL COMMENT '收货地址ID',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_order_no` (`order_no`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='订单表';

-- 订单商品表
CREATE TABLE IF NOT EXISTS `order_item` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `order_id` int unsigned NOT NULL COMMENT '订单ID',
  `product_id` int unsigned NOT NULL COMMENT '商品ID',
  `quantity` int NOT NULL COMMENT '数量',
  `price` decimal(10,2) NOT NULL COMMENT '单价',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_product_id` (`product_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='订单商品表';

-- 插入基础分类数据
INSERT INTO `category` (`name`, `sort`, `status`) VALUES
('蔬菜', 1, 1),
('水果', 2, 1),
('粮油', 3, 1),
('干货', 4, 1),
('特产', 5, 1);

-- 插入示例商品数据
INSERT INTO `product` (`name`, `category_id`, `price`, `stock`, `cover_img`, `description`, `status`, `sort`) VALUES
('新鲜西红柿', 1, 5.99, 100, 'https://example.com/tomato.jpg', '新鲜采摘的西红柿，酸甜可口', 1, 1),
('有机胡萝卜', 1, 3.99, 80, 'https://example.com/carrot.jpg', '有机种植的胡萝卜，营养丰富', 1, 2),
('红富士苹果', 2, 8.99, 50, 'https://example.com/apple.jpg', '新鲜红富士苹果，脆甜多汁', 1, 1),
('香蕉', 2, 4.99, 60, 'https://example.com/banana.jpg', '进口香蕉，香甜软糯', 1, 2),
('东北大米', 3, 29.99, 30, 'https://example.com/rice.jpg', '东北优质大米，粒粒饱满', 1, 1),
('花生油', 3, 45.99, 20, 'https://example.com/oil.jpg', '纯正花生油，香味浓郁', 1, 2); 