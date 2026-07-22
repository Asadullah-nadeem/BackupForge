package logicaltesting

import (
	"github.com/gin-gonic/gin"

	backups_controllers_logical "BackupForge-backend/internal/features/backups/backups/controllers/logical"
	backups_config_logical "BackupForge-backend/internal/features/backups/config/logical"
	"BackupForge-backend/internal/features/databases"
	"BackupForge-backend/internal/features/restores"
	workspaces_controllers "BackupForge-backend/internal/features/workspaces/controllers"
	workspaces_testing "BackupForge-backend/internal/features/workspaces/testing"
)

// CreateTestRouter builds the Gin router wiring the workspace, database, backup
// config, backup and restore controllers used by every logical backup/restore
// test.
func CreateTestRouter() *gin.Engine {
	return workspaces_testing.CreateTestRouter(
		workspaces_controllers.GetWorkspaceController(),
		workspaces_controllers.GetMembershipController(),
		databases.GetDatabaseController(),
		backups_config_logical.GetBackupConfigController(),
		backups_controllers_logical.GetBackupController(),
		restores.GetRestoreController(),
	)
}
