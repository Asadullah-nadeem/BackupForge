package usecases_logical

import (
	usecases_logical_mariadb "BackupForge-backend/internal/features/backups/backups/usecases/logical/mariadb"
	usecases_logical_mongodb "BackupForge-backend/internal/features/backups/backups/usecases/logical/mongodb"
	usecases_logical_mysql "BackupForge-backend/internal/features/backups/backups/usecases/logical/mysql"
	usecases_logical_postgresql "BackupForge-backend/internal/features/backups/backups/usecases/logical/postgresql"
)

var createBackupUsecase = &CreateBackupUsecase{
	usecases_logical_postgresql.GetCreatePostgresqlBackupUsecase(),
	usecases_logical_mysql.GetCreateMysqlBackupUsecase(),
	usecases_logical_mariadb.GetCreateMariadbBackupUsecase(),
	usecases_logical_mongodb.GetCreateMongodbBackupUsecase(),
}

func GetCreateBackupUsecase() *CreateBackupUsecase {
	return createBackupUsecase
}
