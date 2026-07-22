package healthcheck_attempt

import (
	"BackupForge-backend/internal/features/databases"
	healthcheck_config "BackupForge-backend/internal/features/healthcheck/config"
	"BackupForge-backend/internal/features/notifiers"
	workspaces_services "BackupForge-backend/internal/features/workspaces/services"
	"BackupForge-backend/internal/util/logger"
)

var (
	healthcheckAttemptRepository = &HealthcheckAttemptRepository{}
	healthcheckAttemptService    = &HealthcheckAttemptService{
		healthcheckAttemptRepository,
		databases.GetDatabaseService(),
		workspaces_services.GetWorkspaceService(),
	}
)

var checkDatabaseHealthUseCase = &CheckDatabaseHealthUseCase{
	healthcheckAttemptRepository,
	notifiers.GetNotifierService(),
	databases.GetDatabaseService(),
}

var healthcheckAttemptBackgroundService = &HealthcheckAttemptBackgroundService{
	healthcheckConfigService:   healthcheck_config.GetHealthcheckConfigService(),
	checkDatabaseHealthUseCase: checkDatabaseHealthUseCase,
	logger:                     logger.GetLogger(),
}

var healthcheckAttemptController = &HealthcheckAttemptController{
	healthcheckAttemptService,
}

func GetHealthcheckAttemptRepository() *HealthcheckAttemptRepository {
	return healthcheckAttemptRepository
}

func GetHealthcheckAttemptService() *HealthcheckAttemptService {
	return healthcheckAttemptService
}

func GetHealthcheckAttemptBackgroundService() *HealthcheckAttemptBackgroundService {
	return healthcheckAttemptBackgroundService
}

func GetHealthcheckAttemptController() *HealthcheckAttemptController {
	return healthcheckAttemptController
}
