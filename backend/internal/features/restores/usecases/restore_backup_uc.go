package usecases

import (
	"context"
	"errors"

	backups_core_logical "BackupForge-backend/internal/features/backups/backups/core/logical"
	backups_config_logical "BackupForge-backend/internal/features/backups/config/logical"
	"BackupForge-backend/internal/features/databases"
	restores_core "BackupForge-backend/internal/features/restores/core"
	usecases_mariadb "BackupForge-backend/internal/features/restores/usecases/mariadb"
	usecases_mongodb "BackupForge-backend/internal/features/restores/usecases/mongodb"
	usecases_mysql "BackupForge-backend/internal/features/restores/usecases/mysql"
	usecases_postgresql "BackupForge-backend/internal/features/restores/usecases/postgresql"
	"BackupForge-backend/internal/features/storages"
)

type RestoreBackupUsecase struct {
	restorePostgresqlBackupUsecase *usecases_postgresql.RestorePostgresqlBackupUsecase
	restoreMysqlBackupUsecase      *usecases_mysql.RestoreMysqlBackupUsecase
	restoreMariadbBackupUsecase    *usecases_mariadb.RestoreMariadbBackupUsecase
	restoreMongodbBackupUsecase    *usecases_mongodb.RestoreMongodbBackupUsecase
}

func (uc *RestoreBackupUsecase) Execute(
	ctx context.Context,
	backupConfig *backups_config_logical.LogicalBackupConfig,
	restore restores_core.Restore,
	originalDB *databases.Database,
	restoringToDB *databases.Database,
	backup *backups_core_logical.LogicalBackup,
	storage *storages.Storage,
	options restores_core.RestoreOptions,
) error {
	switch originalDB.Type {
	case databases.DatabaseTypePostgresLogical:
		return uc.restorePostgresqlBackupUsecase.Execute(
			ctx,
			originalDB,
			restoringToDB,
			backupConfig,
			restore,
			backup,
			storage,
			options,
		)
	case databases.DatabaseTypeMysql:
		return uc.restoreMysqlBackupUsecase.Execute(
			ctx,
			originalDB,
			restoringToDB,
			backupConfig,
			restore,
			backup,
			storage,
		)
	case databases.DatabaseTypeMariadb:
		return uc.restoreMariadbBackupUsecase.Execute(
			ctx,
			originalDB,
			restoringToDB,
			backupConfig,
			restore,
			backup,
			storage,
		)
	case databases.DatabaseTypeMongodb:
		return uc.restoreMongodbBackupUsecase.Execute(
			ctx,
			originalDB,
			restoringToDB,
			backupConfig,
			restore,
			backup,
			storage,
		)
	default:
		return errors.New("database type not supported")
	}
}
