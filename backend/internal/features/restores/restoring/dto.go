package restoring

import (
	"BackupForge-backend/internal/features/databases/databases/mariadb"
	"BackupForge-backend/internal/features/databases/databases/mongodb"
	"BackupForge-backend/internal/features/databases/databases/mysql"
	postgresql_logical "BackupForge-backend/internal/features/databases/databases/postgresql/logical"
)

type RestoreDatabaseCache struct {
	PostgresqlLogicalDatabase *postgresql_logical.PostgresqlLogicalDatabase `json:"postgresqlDatabase,omitzero"`
	MysqlDatabase             *mysql.MysqlDatabase                          `json:"mysqlDatabase,omitzero"`
	MariadbDatabase           *mariadb.MariadbDatabase                      `json:"mariadbDatabase,omitzero"`
	MongodbDatabase           *mongodb.MongodbDatabase                      `json:"mongodbDatabase,omitzero"`
}
