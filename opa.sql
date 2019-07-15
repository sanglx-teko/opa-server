-- MySQL dump 10.13  Distrib 8.0.15, for osx10.14 (x86_64)
--
-- Host: localhost    Database: opadb
-- ------------------------------------------------------
-- Server version	8.0.15

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8mb4 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `opa_actions`
--

DROP TABLE IF EXISTS `opa_actions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `opa_actions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'read,write',
  `description` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `opa_actions`
--

LOCK TABLES `opa_actions` WRITE;
/*!40000 ALTER TABLE `opa_actions` DISABLE KEYS */;
INSERT INTO `opa_actions` VALUES (1,'read','read','2019-07-10 16:06:58','2019-07-10 16:06:58'),(2,'write','write','2019-07-10 16:06:58','2019-07-10 16:06:58');
/*!40000 ALTER TABLE `opa_actions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `opa_permissions`
--

DROP TABLE IF EXISTS `opa_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `opa_permissions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `action_id` int(11) DEFAULT NULL,
  `resource_id` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `key` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `opa_permissions`
--

LOCK TABLES `opa_permissions` WRITE;
/*!40000 ALTER TABLE `opa_permissions` DISABLE KEYS */;
INSERT INTO `opa_permissions` VALUES (1,1,1,'2019-07-10 16:14:27','2019-07-10 17:28:16','read:users'),(2,2,2,'2019-07-10 16:14:27','2019-07-10 17:28:16','write:permissions');
/*!40000 ALTER TABLE `opa_permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `opa_resources`
--

DROP TABLE IF EXISTS `opa_resources`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `opa_resources` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `service_id` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `opa_resources`
--

LOCK TABLES `opa_resources` WRITE;
/*!40000 ALTER TABLE `opa_resources` DISABLE KEYS */;
INSERT INTO `opa_resources` VALUES (1,'users',1,'2019-07-10 15:55:43','2019-07-10 15:55:43'),(2,'permissions',1,'2019-07-10 15:55:43','2019-07-10 15:55:43'),(3,'orders',2,'2019-07-10 15:55:43','2019-07-10 15:55:43'),(4,'transactions',2,'2019-07-10 15:55:43','2019-07-10 15:55:43');
/*!40000 ALTER TABLE `opa_resources` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `opa_role_permissions`
--

DROP TABLE IF EXISTS `opa_role_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `opa_role_permissions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) DEFAULT NULL,
  `permission_id` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `opa_role_permissions`
--

LOCK TABLES `opa_role_permissions` WRITE;
/*!40000 ALTER TABLE `opa_role_permissions` DISABLE KEYS */;
INSERT INTO `opa_role_permissions` VALUES (1,1,2,'2019-07-10 17:05:05','2019-07-10 17:05:05'),(2,3,1,'2019-07-10 17:05:05','2019-07-10 17:05:05');
/*!40000 ALTER TABLE `opa_role_permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `opa_roles`
--

DROP TABLE IF EXISTS `opa_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `opa_roles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `opa_roles`
--

LOCK TABLES `opa_roles` WRITE;
/*!40000 ALTER TABLE `opa_roles` DISABLE KEYS */;
INSERT INTO `opa_roles` VALUES (1,'developer','2019-07-10 16:41:46','2019-07-10 16:41:46'),(2,'administrator','2019-07-10 16:41:46','2019-07-10 16:41:46'),(3,'user','2019-07-10 16:41:46','2019-07-10 16:44:47'),(4,'super admin','2019-07-10 16:41:46','2019-07-10 16:41:46');
/*!40000 ALTER TABLE `opa_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `opa_service_groups`
--

DROP TABLE IF EXISTS `opa_service_groups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `opa_service_groups` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `uri` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `opa_service_groups`
--

LOCK TABLES `opa_service_groups` WRITE;
/*!40000 ALTER TABLE `opa_service_groups` DISABLE KEYS */;
INSERT INTO `opa_service_groups` VALUES (1,'core_services','http://localhost:8181/v1/data/rbac/authz/allow','2019-07-10 15:04:23','2019-07-11 17:50:24'),(2,'online_sales','http://localhost:8181/v1/data/rbac/authz/allow','2019-07-10 15:04:23','2019-07-11 17:50:23');
/*!40000 ALTER TABLE `opa_service_groups` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `opa_services`
--

DROP TABLE IF EXISTS `opa_services`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `opa_services` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `service_info` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `service_group_id` int(11) DEFAULT NULL,
  `service_metadata` json DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `opa_services`
--

LOCK TABLES `opa_services` WRITE;
/*!40000 ALTER TABLE `opa_services` DISABLE KEYS */;
INSERT INTO `opa_services` VALUES (1,'IAM','Identity API Platform',1,'{}','2019-07-10 15:10:35','2019-07-11 14:11:04'),(2,'PMAPI','Payment API',1,'{}','2019-07-10 15:10:35','2019-07-11 14:11:04');
/*!40000 ALTER TABLE `opa_services` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `opa_user_roles`
--

DROP TABLE IF EXISTS `opa_user_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `opa_user_roles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `opa_user_roles`
--

LOCK TABLES `opa_user_roles` WRITE;
/*!40000 ALTER TABLE `opa_user_roles` DISABLE KEYS */;
INSERT INTO `opa_user_roles` VALUES (1,'U01',1,'2019-07-10 16:54:13','2019-07-10 16:54:13'),(2,'U02',3,'2019-07-10 16:54:13','2019-07-10 16:54:13');
/*!40000 ALTER TABLE `opa_user_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `opa_users`
--

DROP TABLE IF EXISTS `opa_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `opa_users` (
  `id` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `email` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `phone_number` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `address` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `description` text COLLATE utf8mb4_general_ci,
  `birthday` date DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `phone_number` (`phone_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `opa_users`
--

LOCK TABLES `opa_users` WRITE;
/*!40000 ALTER TABLE `opa_users` DISABLE KEYS */;
INSERT INTO `opa_users` VALUES ('U01','Ly Xuan Sang','sang.lxuan@gmail.com','+84347942877','HaNoi','Day la Sang','1991-04-02','2019-07-10 16:49:21','2019-07-10 16:49:21'),('U02','Le Hai Nam','lehainam.dev@gmail.com','','Quang Ninh','Day la Nam','1995-04-02','2019-07-10 16:50:03','2019-07-11 10:58:46');
/*!40000 ALTER TABLE `opa_users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-07-15 11:13:50
