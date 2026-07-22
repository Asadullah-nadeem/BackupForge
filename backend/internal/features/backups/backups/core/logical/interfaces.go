package backups_core_logical

import (
	"context"

	"github.com/google/uuid"

	backups_config_logical "BackupForge-backend/internal/features/backups/config/logical"
	"BackupForge-backend/internal/features/databases"
	"BackupForge-backend/internal/features/notifiers"
	notifier_models "BackupForge-backend/internal/features/notifiers/models"
	"BackupForge-backend/internal/features/storages"
)

type NotificationSender interface {
	SendNotification(
		notifier *notifiers.Notifier,
		notification notifier_models.Notification,
	)
}

type CreateBackupUsecase interface {
	Execute(
		ctx context.Context,
		backup *LogicalBackup,
		backupConfig *backups_config_logical.LogicalBackupConfig,
		database *databases.Database,
		storage *storages.Storage,
		backupProgressListener func(completedMBs float64),
	) (*BackupMetadata, error)
}

type BackupRemoveListener interface {
	OnBeforeBackupRemove(backup *LogicalBackup) error
}

type BackupCompletionListener interface {
	OnBackupCompleted(backupID uuid.UUID)
}
