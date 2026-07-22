package telemetry

import (
	"sync"

	"BackupForge-backend/internal/config"
	physical_service "BackupForge-backend/internal/features/backups/backups/core/physical/service"
	backups_services "BackupForge-backend/internal/features/backups/backups/services"
	"BackupForge-backend/internal/features/databases"
	"BackupForge-backend/internal/features/notifiers"
	"BackupForge-backend/internal/features/storages"
	system_version "BackupForge-backend/internal/features/system/version"
	users_services "BackupForge-backend/internal/features/users/services"
	verification_agents "BackupForge-backend/internal/features/verification/agents"
	verification_config "BackupForge-backend/internal/features/verification/config"
	"BackupForge-backend/internal/util/logger"
)

const productionEndpoint = "https://metrics.BackupForge.com/api/anonymous/collect"

var (
	telemetryLogger = logger.GetLogger()

	instanceLoader = NewInstanceFileLoader("", telemetryLogger)
	httpSender     = NewHTTPTelemetrySender(productionEndpoint, system_version.GetAppVersion())

	telemetryService = NewTelemetryService(
		instanceLoader,
		httpSender,
		databases.GetDatabaseService(),
		storages.GetStorageService(),
		notifiers.GetNotifierService(),
		backups_services.GetBackupService(),
		physical_service.GetPhysicalBackupService(),
		verification_agents.GetAgentService(),
		verification_config.GetVerificationConfigService(),
		users_services.GetUserService(),
		system_version.GetAppVersion(),
		telemetryLogger,
	)

	telemetryBackgroundService = NewTelemetryBackgroundService(
		telemetryService,
		telemetryLogger,
	)
)

func GetTelemetryService() *TelemetryService {
	return telemetryService
}

func GetTelemetryBackgroundService() *TelemetryBackgroundService {
	return telemetryBackgroundService
}

var SetupDependencies = sync.OnceFunc(func() {
	instanceLoader.path = config.GetEnv().TelemetryInstancePath
})
