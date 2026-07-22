package usecases_logical

import (
	"context"
	"errors"

	backups_core_logical "BackupForge-backend/internal/features/backups/backups/core/logical"
	usecases_logical_mariadb "BackupForge-backend/internal/features/backups/backups/usecases/logical/mariadb"
	usecases_logical_mongodb "BackupForge-backend/internal/features/backups/backups/usecases/logical/mongodb"
	usecases_logical_mysql "BackupForge-backend/internal/features/backups/backups/usecases/logical/mysql"
	usecases_logical_postgresql "BackupForge-backend/internal/features/backups/backups/usecases/logical/postgresql"
	backups_config_logical "BackupForge-backend/internal/features/backups/config/logical"
	"BackupForge-backend/internal/features/databases"
	"BackupForge-backend/internal/features/storages"
)

type CreateBackupUsecase struct {
	CreatePostgresqlBackupUsecase *usecases_logical_postgresql.CreatePostgresqlBackupUsecase
	CreateMysqlBackupUsecase      *usecases_logical_mysql.CreateMysqlBackupUsecase
	CreateMariadbBackupUsecase    *usecases_logical_mariadb.CreateMariadbBackupUsecase
	CreateMongodbBackupUsecase    *usecases_logical_mongodb.CreateMongodbBackupUsecase
}

func (uc *CreateBackupUsecase) Execute(
	ctx context.Context,
	backup *backups_core_logical.LogicalBackup,
	backupConfig *backups_config_logical.LogicalBackupConfig,
	database *databases.Database,
	storage *storages.Storage,
	backupProgressListener func(completedMBs float64),
) (*backups_core_logical.BackupMetadata, error) {
	switch database.Type {
	case databases.DatabaseTypePostgresLogical:
		return uc.CreatePostgresqlBackupUsecase.Execute(
			ctx,
			backup,
			backupConfig,
			database,
			storage,
			backupProgressListener,
		)

	case databases.DatabaseTypeMysql:
		return uc.CreateMysqlBackupUsecase.Execute(
			ctx,
			backup,
			backupConfig,
			database,
			storage,
			backupProgressListener,
		)

	case databases.DatabaseTypeMariadb:
		return uc.CreateMariadbBackupUsecase.Execute(
			ctx,
			backup,
			backupConfig,
			database,
			storage,
			backupProgressListener,
		)

	case databases.DatabaseTypeMongodb:
		return uc.CreateMongodbBackupUsecase.Execute(
			ctx,
			backup,
			backupConfig,
			database,
			storage,
			backupProgressListener,
		)

	default:
		return nil, errors.New("database type not supported")
	}
}
