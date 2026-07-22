package system_healthcheck

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	backuping_logical "BackupForge-backend/internal/features/backups/backups/backuping/logical"
	"BackupForge-backend/internal/features/disk"
	verification_agents "BackupForge-backend/internal/features/verification/agents"
	verification_runs "BackupForge-backend/internal/features/verification/runs"
	"BackupForge-backend/internal/storage"
	cache_utils "BackupForge-backend/internal/util/cache"
	"BackupForge-backend/internal/util/tools"
)

type HealthcheckService struct {
	diskService             *disk.DiskService
	backupBackgroundService *backuping_logical.BackupsScheduler
	agentService            *verification_agents.AgentService
}

func (s *HealthcheckService) IsHealthy() error {
	return s.performHealthCheck()
}

func (s *HealthcheckService) performHealthCheck() error {
	// Check if cache is available with PING
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client := cache_utils.GetValkeyClient()
	pingResult := client.Do(ctx, client.B().Ping().Build())
	if pingResult.Error() != nil {
		return errors.New("cannot connect to valkey")
	}

	diskUsage, err := s.diskService.GetDiskUsage()
	if err != nil {
		return errors.New("cannot get disk usage")
	}

	if float64(diskUsage.UsedSpaceBytes) >= float64(diskUsage.TotalSpaceBytes)*0.95 {
		return errors.New("more than 95% of the disk is used")
	}

	if err := tools.ClientToolsHealthError(); err != nil {
		return err
	}

	db := storage.GetDb()
	err = db.Raw("SELECT 1").Error
	if err != nil {
		return errors.New("cannot connect to the database")
	}

	if !s.backupBackgroundService.IsSchedulerRunning() {
		return errors.New("backups are not running for more than 5 minutes")
	}

	staleAgents, err := s.agentService.GetStaleAgents(verification_runs.StaleAgentThreshold)
	if err != nil {
		return errors.New("cannot query verification agents")
	}

	if len(staleAgents) > 0 {
		names := make([]string, len(staleAgents))
		for i, agent := range staleAgents {
			names[i] = agent.Name
		}

		return fmt.Errorf(
			"verification agents not seen for more than 5 minutes: %s",
			strings.Join(names, ", "),
		)
	}

	return nil
}
