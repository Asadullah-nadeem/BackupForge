package system_healthcheck

import (
	backuping_logical "BackupForge-backend/internal/features/backups/backups/backuping/logical"
	"BackupForge-backend/internal/features/disk"
	verification_agents "BackupForge-backend/internal/features/verification/agents"
)

var healthcheckService = &HealthcheckService{
	disk.GetDiskService(),
	backuping_logical.GetBackupsScheduler(),
	verification_agents.GetAgentService(),
}

var healthcheckController = &HealthcheckController{
	healthcheckService,
}

func GetHealthcheckController() *HealthcheckController {
	return healthcheckController
}
