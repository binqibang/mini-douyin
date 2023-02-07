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

 Date: 07/02/2023 21:01:36
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for video_xes
-- ----------------------------
DROP TABLE IF EXISTS `video_xes`;
CREATE TABLE `video_xes`  (
  `id` int NOT NULL,
  `author` varchar(255) CHARACTER SET utf8 COLLATE utf8_croatian_ci NULL DEFAULT NULL,
  `play_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `cover_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `favorite_count` int NULL DEFAULT NULL,
  `comment_count` int NULL DEFAULT NULL,
  `is_favorite` int NULL DEFAULT NULL,
  `publish_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of video_xes
-- ----------------------------
INSERT INTO `video_xes` VALUES (1, '1', '1.mp4', '1', 1, 1, 1, '2023-01-18 11:48:48');
INSERT INTO `video_xes` VALUES (2, '2', '2.mp4', '2', 2, 2, 0, '2023-01-18 11:48:51');
INSERT INTO `video_xes` VALUES (3, '3', '3.mp4', '3', 3, 3, 0, '2023-01-18 11:48:54');
INSERT INTO `video_xes` VALUES (4, '4', '4.mp4', '4', 4, 4, 0, '2023-01-18 11:48:56');
INSERT INTO `video_xes` VALUES (5, '5', '5.mp4', '5', 5, 5, 0, '2023-01-18 11:49:00');
INSERT INTO `video_xes` VALUES (6, '6', '6.mp4', '6', 6, 6, 0, '2023-01-18 11:49:02');
INSERT INTO `video_xes` VALUES (7, '7', '7.mp4', '7', 7, 7, 1, '2023-01-18 11:49:05');
INSERT INTO `video_xes` VALUES (8, '8', '8.mp4', '8', 8, 8, 1, '2023-01-18 11:49:08');
INSERT INTO `video_xes` VALUES (9, '9', '9.mp4', '9', 9, 9, 1, '2023-01-18 11:49:11');
INSERT INTO `video_xes` VALUES (10, '10', '10.mp4', '10', 10, 10, 1, '2023-01-18 11:49:13');
INSERT INTO `video_xes` VALUES (11, '11', '11.mp4', '11', 11, 11, 1, '2023-01-18 13:25:49');
INSERT INTO `video_xes` VALUES (12, '121', '12.mp4', '12', 12, 12, 1, '2023-01-18 13:26:11');

SET FOREIGN_KEY_CHECKS = 1;
