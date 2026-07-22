package restores_core

import (
	"BackupForge-backend/internal/features/databases/databases/mariadb"
	"BackupForge-backend/internal/features/databases/databases/mongodb"
	"BackupForge-backend/internal/features/databases/databases/mysql"
	postgresql_logical "BackupForge-backend/internal/features/databases/databases/postgresql/logical"
)

type RestoreBackupRequest struct {
	PostgresqlLogicalDatabase *postgresql_logical.PostgresqlLogicalDatabase `json:"postgresqlDatabase"`
	MysqlDatabase             *mysql.MysqlDatabase                          `json:"mysqlDatabase"`
	MariadbDatabase           *mariadb.MariadbDatabase                      `json:"mariadbDatabase"`
	MongodbDatabase           *mongodb.MongodbDatabase                      `json:"mongodbDatabase"`
}

type RestoreOptions struct {
	IsExcludeExtensions bool
	IsSkipUserMappings  bool
}
