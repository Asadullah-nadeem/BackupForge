package verification_agents

import (
	audit_logs "BackupForge-backend/internal/features/audit_logs"
	cache_utils "BackupForge-backend/internal/util/cache"
	"BackupForge-backend/internal/util/logger"
)

var agentRepository = &AgentRepository{}

var agentService = &AgentService{
	agentRepository,
	audit_logs.GetAuditLogService(),
	cache_utils.NewRateLimiter(cache_utils.GetValkeyClient()),
	logger.GetLogger(),
	nil,
}

var agentController = &AgentController{
	agentService,
}

var agentFacingController = &AgentFacingController{
	agentService,
}

func GetAgentService() *AgentService {
	return agentService
}

func GetAgentController() *AgentController {
	return agentController
}

func GetAgentFacingController() *AgentFacingController {
	return agentFacingController
}
