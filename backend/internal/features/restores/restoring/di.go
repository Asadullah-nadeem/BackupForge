package restoring

import (
	"sync/atomic"
	"time"

	backups_services "BackupForge-backend/internal/features/backups/backups/services"
	backups_config_logical "BackupForge-backend/internal/features/backups/config/logical"
	"BackupForge-backend/internal/features/databases"
	restores_core "BackupForge-backend/internal/features/restores/core"
	"BackupForge-backend/internal/features/restores/usecases"
	"BackupForge-backend/internal/features/storages"
	tasks_cancellation "BackupForge-backend/internal/features/tasks/cancellation"
	cache_utils "BackupForge-backend/internal/util/cache"
	"BackupForge-backend/internal/util/encryption"
	"BackupForge-backend/internal/util/logger"
)

var restoreRepository = &restores_core.RestoreRepository{}

var restoreDatabaseCache = cache_utils.NewCacheUtil[RestoreDatabaseCache](
	cache_utils.GetValkeyClient(),
	"restore_db:",
)

var restoreCancelManager = tasks_cancellation.GetTaskCancelManager()

var restorer = &Restorer{
	databases.GetDatabaseService(),
	backups_services.GetBackupService(),
	encryption.GetFieldEncryptor(),
	restoreRepository,
	backups_config_logical.GetBackupConfigService(),
	storages.GetStorageService(),
	logger.GetLogger(),
	usecases.GetRestoreBackupUsecase(),
	restoreDatabaseCache,
	restoreCancelManager,
}

var restoresScheduler = &RestoresScheduler{
	restoreRepository,
	time.Now().UTC(),
	logger.GetLogger(),
	restorer,
	restoreDatabaseCache,
	atomic.Bool{},
}

func GetRestoresScheduler() *RestoresScheduler {
	return restoresScheduler
}

func GetRestorer() *Restorer {
	return restorer
}
