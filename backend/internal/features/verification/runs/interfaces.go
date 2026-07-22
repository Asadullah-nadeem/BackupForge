package verification_runs

import (
	"BackupForge-backend/internal/features/notifiers"
	notifier_models "BackupForge-backend/internal/features/notifiers/models"
)

type NotificationSender interface {
	SendNotification(notifier *notifiers.Notifier, notification notifier_models.Notification)
}
