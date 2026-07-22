package verification_config

import (
	"sync"

	"BackupForge-backend/internal/features/audit_logs"
	"BackupForge-backend/internal/features/databases"
	workspaces_services "BackupForge-backend/internal/features/workspaces/services"
	"BackupForge-backend/internal/util/logger"
)

var verificationConfigRepository = &VerificationConfigRepository{}

var verificationConfigService = &VerificationConfigService{
	verificationConfigRepository,
	databases.GetDatabaseService(),
	workspaces_services.GetWorkspaceService(),
	audit_logs.GetAuditLogService(),
	logger.GetLogger(),
}

var verificationConfigController = &VerificationConfigController{
	verificationConfigService,
}

func GetVerificationConfigService() *VerificationConfigService {
	return verificationConfigService
}

func GetVerificationConfigController() *VerificationConfigController {
	return verificationConfigController
}

var SetupDependencies = sync.OnceFunc(func() {
	databases.GetDatabaseService().AddDbCopyListener(verificationConfigService)
})
