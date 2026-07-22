package physical_service

import (
	"BackupForge-backend/internal/features/backups/backups/core/physical/chain_view"
	physical_repositories "BackupForge-backend/internal/features/backups/backups/core/physical/repositories"
	"BackupForge-backend/internal/features/storages"
	"BackupForge-backend/internal/util/encryption"
	"BackupForge-backend/internal/util/logger"
)

var physicalBackupService = &PhysicalBackupService{
	physical_repositories.GetFullBackupRepository(),
	physical_repositories.GetIncrementalBackupRepository(),
	physical_repositories.GetWalSegmentRepository(),
	physical_repositories.GetWalHistoryRepository(),
	chain_view.GetChainViewService(),
	storages.GetStorageService(),
	encryption.GetFieldEncryptor(),
	logger.GetLogger(),
}

func GetPhysicalBackupService() *PhysicalBackupService {
	return physicalBackupService
}
