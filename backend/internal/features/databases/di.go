package databases

import (
	"sync"

	audit_logs "BackupForge-backend/internal/features/audit_logs"
	physical_core_service "BackupForge-backend/internal/features/backups/backups/core/physical/service"
	"BackupForge-backend/internal/features/notifiers"
	users_services "BackupForge-backend/internal/features/users/services"
	workspaces_services "BackupForge-backend/internal/features/workspaces/services"
	"BackupForge-backend/internal/util/encryption"
	"BackupForge-backend/internal/util/logger"
)

var databaseRepository = &DatabaseRepository{}

var databaseService = &DatabaseService{
	databaseRepository,
	notifiers.GetNotifierService(),
	logger.GetLogger(),
	[]DatabaseCreationListener{},
	[]DatabaseRemoveListener{},
	[]DatabaseCopyListener{},
	workspaces_services.GetWorkspaceService(),
	audit_logs.GetAuditLogService(),
	encryption.GetFieldEncryptor(),
	physical_core_service.GetPhysicalBackupService(),
}

var databaseController = &DatabaseController{
	databaseService,
	users_services.GetUserService(),
	workspaces_services.GetWorkspaceService(),
}

func GetDatabaseService() *DatabaseService {
	return databaseService
}

func GetDatabaseController() *DatabaseController {
	return databaseController
}

var SetupDependencies = sync.OnceFunc(func() {
	workspaces_services.GetWorkspaceService().AddWorkspaceDeletionListener(databaseService)
	notifiers.GetNotifierService().SetNotifierDatabaseCounter(databaseService)
})
