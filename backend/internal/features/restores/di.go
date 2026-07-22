package restores

import (
	"sync"

	audit_logs "BackupForge-backend/internal/features/audit_logs"
	backuping_logical "BackupForge-backend/internal/features/backups/backups/backuping/logical"
	backups_services "BackupForge-backend/internal/features/backups/backups/services"
	backups_config_logical "BackupForge-backend/internal/features/backups/config/logical"
	"BackupForge-backend/internal/features/databases"
	"BackupForge-backend/internal/features/disk"
	restores_core "BackupForge-backend/internal/features/restores/core"
	"BackupForge-backend/internal/features/restores/usecases"
	"BackupForge-backend/internal/features/storages"
	tasks_cancellation "BackupForge-backend/internal/features/tasks/cancellation"
	workspaces_services "BackupForge-backend/internal/features/workspaces/services"
	"BackupForge-backend/internal/util/encryption"
	"BackupForge-backend/internal/util/logger"
)

var (
	restoreRepository = &restores_core.RestoreRepository{}
	restoreService    = &RestoreService{
		backups_services.GetBackupService(),
		restoreRepository,
		storages.GetStorageService(),
		backups_config_logical.GetBackupConfigService(),
		usecases.GetRestoreBackupUsecase(),
		databases.GetDatabaseService(),
		logger.GetLogger(),
		workspaces_services.GetWorkspaceService(),
		audit_logs.GetAuditLogService(),
		encryption.GetFieldEncryptor(),
		disk.GetDiskService(),
		tasks_cancellation.GetTaskCancelManager(),
	}
)

var restoreController = &RestoreController{
	restoreService,
}

func GetRestoreController() *RestoreController {
	return restoreController
}

var SetupDependencies = sync.OnceFunc(func() {
	backups_services.GetBackupService().AddBackupRemoveListener(restoreService)
	backuping_logical.GetBackupCleaner().AddBackupRemoveListener(restoreService)
})
