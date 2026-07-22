package healthcheck_config

import (
	"sync"

	"BackupForge-backend/internal/features/audit_logs"
	"BackupForge-backend/internal/features/databases"
	workspaces_services "BackupForge-backend/internal/features/workspaces/services"
	"BackupForge-backend/internal/util/logger"
)

var (
	healthcheckConfigRepository = &HealthcheckConfigRepository{}
	healthcheckConfigService    = &HealthcheckConfigService{
		databases.GetDatabaseService(),
		healthcheckConfigRepository,
		workspaces_services.GetWorkspaceService(),
		audit_logs.GetAuditLogService(),
		logger.GetLogger(),
	}
)

var healthcheckConfigController = &HealthcheckConfigController{
	healthcheckConfigService,
}

func GetHealthcheckConfigService() *HealthcheckConfigService {
	return healthcheckConfigService
}

func GetHealthcheckConfigController() *HealthcheckConfigController {
	return healthcheckConfigController
}

var SetupDependencies = sync.OnceFunc(func() {
	databases.
		GetDatabaseService().
		AddDbCreationListener(healthcheckConfigService)
})
