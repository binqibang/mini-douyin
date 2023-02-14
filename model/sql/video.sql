/*
 Navicat Premium Data Transfer

 Source Server         : tr
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : mini-douyin

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 14/02/2023 20:50:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`  (
  `id` int NOT NULL,
  `author` varchar(255) CHARACTER SET utf8 COLLATE utf8_croatian_ci NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `play_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `cover_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `favorite_count` int NULL DEFAULT NULL,
  `comment_count` int NULL DEFAULT NULL,
  `is_favorite` int NULL DEFAULT NULL,
  `publish_time` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of video
-- ----------------------------
INSERT INTO `video` VALUES (1, '1', '1', '1.mp4', '1', 1, 1, 1, '2023-02-14 17:14:05.000');
INSERT INTO `video` VALUES (2, '2', '2', '2.mp4', '2', 2, 2, 0, '2023-02-14 17:14:08.000');
INSERT INTO `video` VALUES (3, '3', '3', '3.mp4', '3', 3, 3, 0, '2023-02-14 17:14:10.000');
INSERT INTO `video` VALUES (4, '4', '4', '4.mp4', '4', 4, 4, 0, '2023-02-14 17:14:13.000');
INSERT INTO `video` VALUES (5, '5', '5', '5.mp4', '5', 5, 5, 0, '2023-02-14 17:14:16.000');
INSERT INTO `video` VALUES (6, '6', '6', '6.mp4', '6', 6, 6, 0, '2023-02-14 17:14:20.000');
INSERT INTO `video` VALUES (7, '7', '7', '7.mp4', '7', 7, 7, 1, '2023-02-14 17:14:22.000');
INSERT INTO `video` VALUES (8, '8', '8', '8.mp4', '8', 8, 8, 1, '2023-02-14 17:14:24.000');
INSERT INTO `video` VALUES (9, '9', '9', '9.mp4', '9', 9, 9, 1, '2023-02-14 17:14:26.000');
INSERT INTO `video` VALUES (10, '10', '10', '10.mp4', '10', 10, 10, 1, '2023-02-14 17:14:28.000');
INSERT INTO `video` VALUES (11, '11', '11', '11.mp4', '11', 11, 11, 1, '2023-02-14 17:14:30.000');
INSERT INTO `video` VALUES (12, '121', '121', '12.mp4', '12', 12, 12, 1, '2023-02-14 17:14:32.000');

SET FOREIGN_KEY_CHECKS = 1;
