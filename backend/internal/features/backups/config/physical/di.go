package backups_config_physical

import (
	"sync"

	"BackupForge-backend/internal/features/databases"
	"BackupForge-backend/internal/features/notifiers"
	"BackupForge-backend/internal/features/storages"
	workspaces_services "BackupForge-backend/internal/features/workspaces/services"
)

var (
	backupConfigRepository = &BackupConfigRepository{}
	backupConfigService    = &BackupConfigService{
		backupConfigRepository,
		databases.GetDatabaseService(),
		storages.GetStorageService(),
		notifiers.GetNotifierService(),
		workspaces_services.GetWorkspaceService(),
		nil,
		nil,
	}
)

var backupConfigController = &BackupConfigController{
	backupConfigService,
}

func GetBackupConfigController() *BackupConfigController {
	return backupConfigController
}

func GetBackupConfigService() *BackupConfigService {
	return backupConfigService
}

var SetupDependencies = sync.OnceFunc(func() {
	storages.GetStorageService().AddStorageDatabaseCounter(backupConfigService)
})
