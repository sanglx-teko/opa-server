package test

import (
	"fmt"

	"github.com/sanglx-teko/opa-server/models"
	"github.com/sanglx-teko/opa-server/models/dao"
)

// InitRoleTable ...
func InitRoleTable() {
	db := dao.ConfigurationManager.GetDB()
	qDropTable := fmt.Sprintf("DROP TABLE IF EXISTS `%s`", models.TableOPARole)
	db.Exec(qDropTable)
	qCreateTable := fmt.Sprintf("CREATE TABLE `%s` ("+
		"`id` int(11) NOT NULL AUTO_INCREMENT,"+
		"`name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,"+
		"`created_at` datetime DEFAULT CURRENT_TIMESTAMP,"+
		"`updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,"+
		"PRIMARY KEY (`id`),"+
		"UNIQUE KEY `name` (`name`)"+
		") ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci",
		models.TableOPARole)
	db.Exec(qCreateTable)
	qInsert := fmt.Sprintf("INSERT INTO `%s` VALUES "+
		"(1,'%s','2019-07-10 16:41:46','2019-07-10 16:41:46'),"+
		"(2,'%s','2019-07-10 16:41:46','2019-07-10 16:41:46'),"+
		"(3,'%s','2019-07-10 16:41:46','2019-07-10 16:44:47'),"+
		"(4,'%s','2019-07-10 16:41:46','2019-07-10 16:41:46')",
		models.TableOPARole,
		"developer",
		"administrator",
		"user",
		"super admin")
	db.Exec(qInsert)
}

// InitUserRoleTable ...
func InitUserRoleTable() {
	db := dao.ConfigurationManager.GetDB()
	qDropTable := fmt.Sprintf("DROP TABLE IF EXISTS `%s`", models.TableOPARUserRole)
	db.Exec(qDropTable)
	qCreateTable := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s`("+
		"`id` int(11) NOT NULL AUTO_INCREMENT,"+
		"`user_id` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,"+
		"`role_id` int(11) DEFAULT NULL,"+
		"`created_at` datetime DEFAULT CURRENT_TIMESTAMP,"+
		"`updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,"+
		"PRIMARY KEY (`id`)"+
		") ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci",
		models.TableOPARUserRole)
	db.Exec(qCreateTable)
	qInsert := "INSERT INTO `opa_user_roles` VALUES " +
		"(1,'U01',1,'2019-07-10 16:54:13','2019-07-10 16:54:13')," +
		"(2,'U02',3,'2019-07-10 16:54:13','2019-07-10 16:54:13');"
	db.Exec(qInsert)

}

// InitActionTable ...
func InitActionTable() {
	db := dao.ConfigurationManager.GetDB()
	qDropTable := fmt.Sprintf("DROP TABLE IF EXISTS `%s`", models.TableOPAActions)
	db.Exec(qDropTable)
	qCreateTable := fmt.Sprintf("CREATE TABLE `%s` ("+
		"`id` int(11) NOT NULL AUTO_INCREMENT,"+
		"`name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'read,write',"+
		"`description` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,"+
		"`created_at` datetime DEFAULT CURRENT_TIMESTAMP,"+
		"`updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,"+
		"PRIMARY KEY (`id`)"+
		") ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci",
		models.TableOPAActions)
	db.Exec(qCreateTable)
	qInsert := fmt.Sprintf("INSERT INTO `%s` VALUES "+
		"(1,'read','read','2019-07-10 16:06:58','2019-07-10 16:06:58'),"+
		"(2,'write','write','2019-07-10 16:06:58','2019-07-10 16:06:58')",
		models.TableOPAActions)
	db.Exec(qInsert)
}

// InitPermissionTable ...
func InitPermissionTable() {
	db := dao.ConfigurationManager.GetDB()
	qDropTable := fmt.Sprintf("DROP TABLE IF EXISTS `%s`", models.TableOPAPermission)
	db.Exec(qDropTable)
	qCreateTable := fmt.Sprintf("CREATE TABLE `%s` ("+
		"`id` int(11) NOT NULL AUTO_INCREMENT,"+
		"`action_id` int(11) DEFAULT NULL,"+
		"`resource_id` int(11) DEFAULT NULL,"+
		"`created_at` datetime DEFAULT CURRENT_TIMESTAMP,"+
		"`updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,"+
		"`key` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',"+
		"PRIMARY KEY (`id`)"+
		") ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci",
		models.TableOPAPermission)
	db.Exec(qCreateTable)
	qInsert := fmt.Sprintf("INSERT INTO `%s` VALUES "+
		"(1,1,1,'2019-07-10 16:14:27','2019-07-10 17:28:16','read:users'),"+
		"(2,2,2,'2019-07-10 16:14:27','2019-07-10 17:28:16','write:permissions')",
		models.TableOPAPermission)
	db.Exec(qInsert)
}

// InitResourceTable ...
func InitResourceTable() {
	db := dao.ConfigurationManager.GetDB()
	qDropTable := fmt.Sprintf("DROP TABLE IF EXISTS `%s`", models.TableOPAResource)
	db.Exec(qDropTable)
	qCreateTable := fmt.Sprintf("CREATE TABLE `%s` ("+
		"`id` int(11) NOT NULL AUTO_INCREMENT,"+
		"`name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,"+
		"`service_id` int(11) DEFAULT NULL,"+
		"`created_at` datetime DEFAULT CURRENT_TIMESTAMP,"+
		"`updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,"+
		"PRIMARY KEY (`id`),"+
		"UNIQUE KEY `name` (`name`)"+
		") ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci",
		models.TableOPAResource)
	db.Exec(qCreateTable)
	qInsert := fmt.Sprintf("INSERT INTO `%s` VALUES "+
		"(1,'users',1,'2019-07-10 15:55:43','2019-07-10 15:55:43'),"+
		"(2,'permissions',1,'2019-07-10 15:55:43','2019-07-10 15:55:43'),"+
		"(3,'orders',2,'2019-07-10 15:55:43','2019-07-10 15:55:43'),"+
		"(4,'transactions',2,'2019-07-10 15:55:43','2019-07-10 15:55:43')",
		models.TableOPAResource)
	db.Exec(qInsert)
}

// InitRolePermissionTable ...
func InitRolePermissionTable() {
	db := dao.ConfigurationManager.GetDB()
	qDropTable := fmt.Sprintf("DROP TABLE IF EXISTS `%s`", models.TableOPARolePermissions)
	db.Exec(qDropTable)
	qCreateTable := fmt.Sprintf("CREATE TABLE `%s` ("+
		"`id` int(11) NOT NULL AUTO_INCREMENT,"+
		"`role_id` int(11) DEFAULT NULL,"+
		"`permission_id` int(11) DEFAULT NULL,"+
		"`created_at` datetime DEFAULT CURRENT_TIMESTAMP,"+
		"`updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,"+
		"PRIMARY KEY (`id`)"+
		") ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci",
		models.TableOPARolePermissions)
	db.Exec(qCreateTable)
	qInsert := fmt.Sprintf("INSERT INTO `%s` VALUES "+
		"(1,1,2,'2019-07-10 17:05:05','2019-07-10 17:05:05'),"+
		"(2,3,1,'2019-07-10 17:05:05','2019-07-10 17:05:05')",
		models.TableOPARolePermissions)
	db.Exec(qInsert)
}

// InitServiceGroupTable ...
func InitServiceGroupTable() {
	db := dao.ConfigurationManager.GetDB()
	qDropTable := fmt.Sprintf("DROP TABLE IF EXISTS `%s`", models.TableOPAServiceGroup)
	db.Exec(qDropTable)
	qCreateTable := fmt.Sprintf("CREATE TABLE `%s` ("+
		"`id` int(11) NOT NULL AUTO_INCREMENT,"+
		"`name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,"+
		"`uri` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,"+
		"`created_at` datetime DEFAULT CURRENT_TIMESTAMP,"+
		"`updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,"+
		"PRIMARY KEY (`id`),"+
		"UNIQUE KEY `name` (`name`)"+
		") ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci",
		models.TableOPAServiceGroup)
	db.Exec(qCreateTable)
	qInsert := fmt.Sprintf("INSERT INTO `%s` VALUES "+
		"(1,'core_services','http://localhost:8181/v1/data/rbac/authz/allow','2019-07-10 15:04:23','2019-07-11 17:50:24'),"+
		"(2,'online_sales','http://localhost:8181/v1/data/rbac/authz/allow','2019-07-10 15:04:23','2019-07-11 17:50:23')",
		models.TableOPAServiceGroup)
	db.Exec(qInsert)
}

// InitServiceTable ...
func InitServiceTable() {
	db := dao.ConfigurationManager.GetDB()
	qDropTable := fmt.Sprintf("DROP TABLE IF EXISTS `%s`", models.TableOPAService)
	db.Exec(qDropTable)
	qCreateTable := fmt.Sprintf("CREATE TABLE `%s` ("+
		"`id` int(11) NOT NULL AUTO_INCREMENT,"+
		"`name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,"+
		"`service_info` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,"+
		"`service_group_id` int(11) DEFAULT NULL,"+
		"`service_metadata` json DEFAULT NULL,"+
		"`created_at` datetime DEFAULT CURRENT_TIMESTAMP,"+
		"`updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,"+
		"PRIMARY KEY (`id`),"+
		"UNIQUE KEY `name` (`name`)"+
		") ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci",
		models.TableOPAService)
	db.Exec(qCreateTable)
	qInsert := fmt.Sprintf("INSERT INTO `%s` VALUES "+
		"(1,'IAM','Identity API Platform',1,'{}','2019-07-10 15:10:35','2019-07-11 14:11:04'),"+
		"(2,'PMAPI','Payment API',1,'{}','2019-07-10 15:10:35','2019-07-11 14:11:04')",
		models.TableOPAService)
	db.Exec(qInsert)
}

// InitUserTable ...
func InitUserTable() {
	db := dao.ConfigurationManager.GetDB()
	qDropTable := fmt.Sprintf("DROP TABLE IF EXISTS `%s`", models.TableOPAUser)
	db.Exec(qDropTable)
	qCreateTable := fmt.Sprintf("CREATE TABLE `%s` ("+
		"`id` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,"+
		"`name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,"+
		"`email` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,"+
		"`phone_number` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,"+
		"`address` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,"+
		"`description` text COLLATE utf8mb4_general_ci,"+
		"`birthday` date DEFAULT NULL,"+
		"`created_at` datetime DEFAULT CURRENT_TIMESTAMP,"+
		"`updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,"+
		"PRIMARY KEY (`id`),"+
		"UNIQUE KEY `email` (`email`),"+
		"UNIQUE KEY `phone_number` (`phone_number`)"+
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci",
		models.TableOPAUser)
	db.Exec(qCreateTable)
	qInsert := fmt.Sprintf("INSERT INTO `%s` VALUES "+
		"('U01','Ly Xuan Sang','sang.lxuan@gmail.com','+84347942877','HaNoi','Day la Sang','1991-04-02','2019-07-10 16:49:21','2019-07-10 16:49:21'),"+
		"('U02','Le Hai Nam','lehainam.dev@gmail.com','','Quang Ninh','Day la Nam','1995-04-02','2019-07-10 16:50:03','2019-07-11 10:58:46')",
		models.TableOPAUser)
	db.Exec(qInsert)
}

// InitFullDB ...
func InitFullDB() {
	InitUserTable()
	InitRoleTable()
	InitActionTable()
	InitServiceTable()
	InitResourceTable()
	InitUserRoleTable()
	InitPermissionTable()
	InitServiceGroupTable()
	InitRolePermissionTable()
}

// DropAllDB ...
func DropAllDB() {
	db := dao.ConfigurationManager.GetDB()
	qDropTable := "DROP TABLE IF EXISTS `%s`"
	db.Exec(fmt.Sprintf(qDropTable, models.TableOPARole))
	db.Exec(fmt.Sprintf(qDropTable, models.TableOPARUserRole))
	db.Exec(fmt.Sprintf(qDropTable, models.TableOPAActions))
	db.Exec(fmt.Sprintf(qDropTable, models.TableOPARolePermissions))
	db.Exec(fmt.Sprintf(qDropTable, models.TableOPAResource))
	db.Exec(fmt.Sprintf(qDropTable, models.TableOPAPermission))
	db.Exec(fmt.Sprintf(qDropTable, models.TableOPAServiceGroup))
	db.Exec(fmt.Sprintf(qDropTable, models.TableOPAService))
	db.Exec(fmt.Sprintf(qDropTable, models.TableOPAUser))
}
