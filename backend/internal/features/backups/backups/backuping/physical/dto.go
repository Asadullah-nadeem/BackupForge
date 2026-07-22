package backuping_physical

import (
	"time"

	"BackupForge-backend/internal/features/backups/backups/core/physical/chain_view"
	backups_config_physical "BackupForge-backend/internal/features/backups/config/physical"
	"BackupForge-backend/internal/features/databases"
	"BackupForge-backend/internal/features/storages"
)

type backupContext struct {
	Config    *backups_config_physical.PhysicalBackupConfig
	Database  *databases.Database
	Storage   *storages.Storage
	MasterKey string
}

// chainCandidate pairs a non-extendable chain with its end timestamp so passes
// can order by recency without recomputing it.
type chainCandidate struct {
	view  *chain_view.ChainView
	endTs time.Time
}
