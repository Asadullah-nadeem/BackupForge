package verification_runs

import (
	"sync"
	"sync/atomic"

	"BackupForge-backend/internal/features/audit_logs"
	backuping_logical "BackupForge-backend/internal/features/backups/backups/backuping/logical"
	backups_services "BackupForge-backend/internal/features/backups/backups/services"
	"BackupForge-backend/internal/features/databases"
	"BackupForge-backend/internal/features/notifiers"
	verification_agents "BackupForge-backend/internal/features/verification/agents"
	verification_config "BackupForge-backend/internal/features/verification/config"
	workspaces_services "BackupForge-backend/internal/features/workspaces/services"
	"BackupForge-backend/internal/util/logger"
)

var verificationRepository = &VerificationRepository{}

var verificationService = &VerificationService{
	verificationRepository,
	databases.GetDatabaseService(),
	backups_services.GetBackupService(),
	verification_config.GetVerificationConfigService(),
	notifiers.GetNotifierService(),
	workspaces_services.GetWorkspaceService(),
	audit_logs.GetAuditLogService(),
	logger.GetLogger(),
}

var verificationScheduler = &VerificationScheduler{
	verificationRepository,
	verificationService,
	verification_config.GetVerificationConfigService(),
	verification_agents.GetAgentService(),
	backups_services.GetBackupService(),
	databases.GetDatabaseService(),
	logger.GetLogger(),
	atomic.Bool{},
}

var verificationController = &VerificationController{
	verificationService,
}

var verificationAgentController = &VerificationAgentController{
	verificationService,
	verification_agents.GetAgentService(),
	logger.GetLogger(),
}

func GetVerificationService() *VerificationService {
	return verificationService
}

func GetVerificationScheduler() *VerificationScheduler {
	return verificationScheduler
}

func GetVerificationController() *VerificationController {
	return verificationController
}

func GetVerificationAgentController() *VerificationAgentController {
	return verificationAgentController
}

var SetupDependencies = sync.OnceFunc(func() {
	verification_agents.GetAgentService().AddAgentHeartbeatedListener(verificationService)
	backuping_logical.GetBackupsScheduler().AddBackupCompletionListener(verificationService)
})
