/*
 Navicat Premium Data Transfer

 Source Server         : blogstudy
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 01/05/2021 20:18:13
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `cid` bigint(20) UNSIGNED NOT NULL,
  `desc` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `content` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `img` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `comment_count` bigint(20) NOT NULL DEFAULT 0,
  `read_count` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_article_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 15 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES (1, '2021-04-06 14:51:40.625', '2021-04-06 14:51:40.625', NULL, '第一篇', 2, '第一篇描述', '第一篇内容', 'http://qri5ziugf.hd-bkt.clouddn.com/FvFCbPz8i8nVgikPX3fltH1T9udJ', 0, 1);
INSERT INTO `article` VALUES (2, '2021-04-06 14:51:51.077', '2021-04-06 14:51:51.077', NULL, '第2篇', 2, '第2篇描述', '第2篇内容', '第2篇图片', 0, 1);
INSERT INTO `article` VALUES (3, '2021-04-06 14:52:10.981', '2021-04-06 14:52:10.981', NULL, '第3篇', 2, '第3篇描述', '第3篇内容', '第3篇图片', 0, 1);
INSERT INTO `article` VALUES (4, '2021-04-06 14:59:33.811', '2021-04-06 16:02:20.805', '2021-04-06 16:02:52.369', '第4篇', 3, '第4篇描述', '第4篇内容', '第4篇图片', 0, 1);
INSERT INTO `article` VALUES (7, '2021-04-17 22:16:31.197', '2021-04-17 22:16:31.197', '2021-04-17 23:14:11.588', '新增', 7, '新增新增新增', '新增新增新增', '', 0, 0);
INSERT INTO `article` VALUES (5, '2021-04-06 16:03:03.804', '2021-04-06 16:03:03.804', NULL, '第5篇', 3, '第5篇描述', '第5篇内容', '第5篇图片', 0, 1);
INSERT INTO `article` VALUES (6, '2021-04-09 18:12:34.336', '2021-04-21 21:21:17.853', NULL, '第5篇123', 8, '第5篇描述123', '第5篇内容123', '第5篇图片', 0, 0);
INSERT INTO `article` VALUES (8, '2021-04-17 22:16:33.227', '2021-04-17 22:16:33.227', '2021-04-17 23:14:09.957', '新增', 7, '新增新增新增', '新增新增新增', '', 0, 0);
INSERT INTO `article` VALUES (9, '2021-04-17 22:16:33.595', '2021-04-17 22:16:33.595', '2021-04-17 23:14:08.933', '新增', 7, '新增新增新增', '新增新增新增', '', 0, 0);
INSERT INTO `article` VALUES (10, '2021-04-17 22:16:33.789', '2021-04-17 22:16:33.789', '2021-04-17 23:14:08.148', '新增', 7, '新增新增新增', '新增新增新增', '', 0, 0);
INSERT INTO `article` VALUES (11, '2021-04-17 22:16:33.970', '2021-04-17 22:16:33.970', '2021-04-17 23:14:07.291', '新增', 7, '新增新增新增', '新增新增新增', '', 0, 0);
INSERT INTO `article` VALUES (12, '2021-04-17 23:09:33.369', '2021-04-17 23:10:16.865', '2021-04-17 23:14:06.374', '新增新增新增新增', 7, '新增新增新增', '新增新增新增', '', 0, 0);
INSERT INTO `article` VALUES (13, '2021-04-17 23:11:38.383', '2021-04-17 23:11:38.383', '2021-04-17 23:14:05.229', '/admin/artlist', 2, '/admin/artlist', '/admin/artlist/admin/artlist', '', 0, 0);
INSERT INTO `article` VALUES (14, '2021-04-21 21:02:03.428', '2021-04-21 21:02:03.428', NULL, '新增', 2, '新增新增新增', '新增新增新增新增新增', 'http://cdn.tsuhao.tech/Fq9pZdQg5pvpMpiQd78rChv2w5di', 0, 0);

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (7, 'testweb1');
INSERT INTO `category` VALUES (2, 'testedit');
INSERT INTO `category` VALUES (3, 'test3');
INSERT INTO `category` VALUES (4, 'test4');
INSERT INTO `category` VALUES (5, 'test5');
INSERT INTO `category` VALUES (9, '123');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `username` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `role` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 20 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '2021-04-14 20:23:04.910', '2021-04-14 20:23:04.910', NULL, 'admin', 'zB5vqQ86HphjoQ==', 1);
INSERT INTO `user` VALUES (2, '2021-04-14 20:40:56.923', '2021-04-15 17:37:54.003', '2021-04-15 19:22:21.876', 'user1', '', 2);
INSERT INTO `user` VALUES (3, '2021-04-15 13:40:45.607', '2021-04-16 00:25:41.179', NULL, 'user2edit', 'zB5vqQ86HphjoQ==', 2);
INSERT INTO `user` VALUES (4, '2021-04-15 13:40:47.844', '2021-04-15 13:40:47.844', NULL, 'user3', 'zB5vqQ86HphjoQ==', 2);
INSERT INTO `user` VALUES (5, '2021-04-15 13:40:52.850', '2021-04-15 13:40:52.850', NULL, 'user4', 'zB5vqQ86HphjoQ==', 2);
INSERT INTO `user` VALUES (6, '2021-04-15 13:40:57.383', '2021-04-15 13:40:57.383', NULL, 'user5', 'zB5vqQ86HphjoQ==', 2);
INSERT INTO `user` VALUES (7, '2021-04-15 13:41:00.213', '2021-04-15 13:41:00.213', NULL, 'user6', 'zB5vqQ86HphjoQ==', 2);
INSERT INTO `user` VALUES (8, '2021-04-15 13:41:03.460', '2021-04-15 13:41:03.460', NULL, 'user7', 'zB5vqQ86HphjoQ==', 2);
INSERT INTO `user` VALUES (9, '2021-04-15 13:41:12.169', '2021-04-15 13:41:12.169', NULL, 'admin2', 'zB5vqQ86HphjoQ==', 1);
INSERT INTO `user` VALUES (10, '2021-04-15 13:41:15.144', '2021-04-15 13:41:15.144', NULL, 'admin3', 'zB5vqQ86HphjoQ==', 1);
INSERT INTO `user` VALUES (11, '2021-04-15 15:34:32.150', '2021-04-15 15:34:32.150', NULL, 'user8', 'zB5vqQ86HphjoQ==', 2);
INSERT INTO `user` VALUES (17, '2021-04-15 21:52:28.307', '2021-04-15 21:52:28.307', NULL, 'user1', 'RsDR3DP0ZysLYw==', 2);
INSERT INTO `user` VALUES (16, '2021-04-15 21:46:39.187', '2021-04-15 21:46:39.187', NULL, '123123', 'RsDR3DP0ZysLYw==', 2);
INSERT INTO `user` VALUES (15, '2021-04-15 21:41:58.178', '2021-04-15 21:41:58.178', '2021-04-15 21:42:28.335', 'user1', 'zB5vqQ86HphjoQ==', 2);
INSERT INTO `user` VALUES (18, '2021-04-15 21:56:19.313', '2021-04-16 00:50:15.886', '2021-04-16 00:51:51.145', '', 'RsDR3DP0ZysLYw==', 0);
INSERT INTO `user` VALUES (19, '2021-04-16 00:44:00.464', '2021-04-16 00:44:00.464', NULL, '12312312312', 'RsDR3DP0ZysLYw==', 2);

SET FOREIGN_KEY_CHECKS = 1;
