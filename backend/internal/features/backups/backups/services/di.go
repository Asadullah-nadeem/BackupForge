package backups_services

import (
	"sync"

	audit_logs "BackupForge-backend/internal/features/audit_logs"
	backuping_logical "BackupForge-backend/internal/features/backups/backups/backuping/logical"
	backuping_physical "BackupForge-backend/internal/features/backups/backups/backuping/physical"
	backups_core_logical "BackupForge-backend/internal/features/backups/backups/core/logical"
	"BackupForge-backend/internal/features/backups/backups/core/physical/chain_view"
	physical_repositories "BackupForge-backend/internal/features/backups/backups/core/physical/repositories"
	physical_core_service "BackupForge-backend/internal/features/backups/backups/core/physical/service"
	backups_download "BackupForge-backend/internal/features/backups/backups/download"
	usecases_logical "BackupForge-backend/internal/features/backups/backups/usecases/logical"
	backups_config_logical "BackupForge-backend/internal/features/backups/config/logical"
	backups_config_physical "BackupForge-backend/internal/features/backups/config/physical"
	"BackupForge-backend/internal/features/databases"
	encryption_secrets "BackupForge-backend/internal/features/encryption/secrets"
	"BackupForge-backend/internal/features/notifiers"
	"BackupForge-backend/internal/features/storages"
	task_cancellation "BackupForge-backend/internal/features/tasks/cancellation"
	workspaces_services "BackupForge-backend/internal/features/workspaces/services"
	"BackupForge-backend/internal/util/encryption"
	"BackupForge-backend/internal/util/logger"
)

var taskCancelManager = task_cancellation.GetTaskCancelManager()

var backupService = &LogicalBackupService{
	databases.GetDatabaseService(),
	storages.GetStorageService(),
	backups_core_logical.GetBackupRepository(),
	notifiers.GetNotifierService(),
	notifiers.GetNotifierService(),
	backups_config_logical.GetBackupConfigService(),
	encryption_secrets.GetSecretKeyService(),
	encryption.GetFieldEncryptor(),
	usecases_logical.GetCreateBackupUsecase(),
	logger.GetLogger(),
	[]backups_core_logical.BackupRemoveListener{},
	workspaces_services.GetWorkspaceService(),
	audit_logs.GetAuditLogService(),
	taskCancelManager,
	backups_download.GetDownloadTokenService(),
	backuping_logical.GetBackupsScheduler(),
	backuping_logical.GetBackupCleaner(),
}

func GetBackupService() *LogicalBackupService {
	return backupService
}

var physicalBackupService = &PhysicalBackupService{
	databases.GetDatabaseService(),
	workspaces_services.GetWorkspaceService(),
	backups_config_physical.GetBackupConfigService(),
	chain_view.GetChainViewService(),
	physical_core_service.GetPhysicalBackupService(),
	physical_repositories.GetFullBackupRepository(),
	physical_repositories.GetIncrementalBackupRepository(),
	physical_repositories.GetWalSegmentRepository(),
	physical_repositories.GetInFlightBackupRepository(),
	backuping_physical.GetPhysicalBackupCanceller(),
	backups_download.GetRestoreStreamWriter(),
	backups_download.GetRestoreTokenService(),
	encryption_secrets.GetSecretKeyService(),
	audit_logs.GetAuditLogService(),
	logger.GetLogger(),
}

func GetPhysicalBackupService() *PhysicalBackupService {
	return physicalBackupService
}

var SetupDependencies = sync.OnceFunc(func() {
	backups_config_logical.
		GetBackupConfigService().
		SetDatabaseStorageChangeListener(backupService)

	databases.GetDatabaseService().AddDbRemoveListener(backupService)
	databases.GetDatabaseService().AddDbCopyListener(backups_config_logical.GetBackupConfigService())
})
