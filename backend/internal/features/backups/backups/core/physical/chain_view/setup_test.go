package chain_view_test

import (
	"os"
	"testing"

	backuping_physical "BackupForge-backend/internal/features/backups/backups/backuping/physical"
	cache_utils "BackupForge-backend/internal/util/cache"
)

func TestMain(m *testing.M) {
	cache_utils.ClearAllCache()
	backuping_physical.SetupDependencies()

	exitCode := m.Run()

	os.Exit(exitCode)
}
