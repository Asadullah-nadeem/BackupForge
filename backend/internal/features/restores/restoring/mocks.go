package restoring

import (
	"context"
	"errors"

	backups_core_logical "BackupForge-backend/internal/features/backups/backups/core/logical"
	backups_config_logical "BackupForge-backend/internal/features/backups/config/logical"
	"BackupForge-backend/internal/features/databases"
	restores_core "BackupForge-backend/internal/features/restores/core"
	"BackupForge-backend/internal/features/storages"
)

type MockSuccessRestoreUsecase struct{}

func (uc *MockSuccessRestoreUsecase) Execute(
	ctx context.Context,
	backupConfig *backups_config_logical.LogicalBackupConfig,
	restore restores_core.Restore,
	originalDB *databases.Database,
	restoringToDB *databases.Database,
	backup *backups_core_logical.LogicalBackup,
	storage *storages.Storage,
	options restores_core.RestoreOptions,
) error {
	return nil
}

type MockFailedRestoreUsecase struct{}

func (uc *MockFailedRestoreUsecase) Execute(
	ctx context.Context,
	backupConfig *backups_config_logical.LogicalBackupConfig,
	restore restores_core.Restore,
	originalDB *databases.Database,
	restoringToDB *databases.Database,
	backup *backups_core_logical.LogicalBackup,
	storage *storages.Storage,
	options restores_core.RestoreOptions,
) error {
	return errors.New("restore failed")
}

type MockCaptureCredentialsRestoreUsecase struct {
	CalledChan    chan *databases.Database
	ShouldSucceed bool
}

func (uc *MockCaptureCredentialsRestoreUsecase) Execute(
	ctx context.Context,
	backupConfig *backups_config_logical.LogicalBackupConfig,
	restore restores_core.Restore,
	originalDB *databases.Database,
	restoringToDB *databases.Database,
	backup *backups_core_logical.LogicalBackup,
	storage *storages.Storage,
	options restores_core.RestoreOptions,
) error {
	uc.CalledChan <- restoringToDB

	if uc.ShouldSucceed {
		return nil
	}
	return errors.New("mock restore failed")
}

type MockBlockingRestoreUsecase struct {
	StartedChan chan bool
}

func (uc *MockBlockingRestoreUsecase) Execute(
	ctx context.Context,
	backupConfig *backups_config_logical.LogicalBackupConfig,
	restore restores_core.Restore,
	originalDB *databases.Database,
	restoringToDB *databases.Database,
	backup *backups_core_logical.LogicalBackup,
	storage *storages.Storage,
	options restores_core.RestoreOptions,
) error {
	if uc.StartedChan != nil {
		uc.StartedChan <- true
	}

	<-ctx.Done()

	return ctx.Err()
}
