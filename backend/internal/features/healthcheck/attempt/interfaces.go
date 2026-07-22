package healthcheck_attempt

import (
	"github.com/google/uuid"

	"BackupForge-backend/internal/features/databases"
	"BackupForge-backend/internal/features/notifiers"
	notifier_models "BackupForge-backend/internal/features/notifiers/models"
)

type HealthcheckAttemptSender interface {
	SendNotification(
		notifier *notifiers.Notifier,
		notification notifier_models.Notification,
	)
}

type DatabaseService interface {
	GetDatabaseByID(id uuid.UUID) (*databases.Database, error)

	TestDatabaseConnectionDirect(database *databases.Database) error

	SetHealthStatus(
		databaseID uuid.UUID,
		healthStatus *databases.HealthStatus,
	) error
}
