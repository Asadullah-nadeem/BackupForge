package backups_download

import (
	"BackupForge-backend/internal/features/backups/backups/download/download_token"
	"BackupForge-backend/internal/features/backups/backups/download/restore_stream"
	"BackupForge-backend/internal/features/backups/backups/download/restore_token"
	"BackupForge-backend/internal/features/backups/backups/download/stream_guard"
	"BackupForge-backend/internal/features/storages"
	cache_utils "BackupForge-backend/internal/util/cache"
	util_encryption "BackupForge-backend/internal/util/encryption"
	"BackupForge-backend/internal/util/logger"
)

// This package is the composition root for the download/restore stream feature.
// The single stream_guard.Guard built here is injected into BOTH token services
// so the per-user single-stream lock is shared — a user can never run a logical
// download and a physical restore stream at the same time.
var (
	downloadTokenService           *download_token.Service
	restoreTokenService            *restore_token.Service
	downloadTokenBackgroundService *download_token.BackgroundService
	restoreStreamWriter            *restore_stream.Writer
)

func init() {
	valkeyClient := cache_utils.GetValkeyClient()

	guard := stream_guard.NewGuard(
		stream_guard.NewTracker(valkeyClient),
		logger.GetLogger(),
	)

	downloadTokenService = download_token.NewService(guard, logger.GetLogger())
	restoreTokenService = restore_token.NewService(guard, valkeyClient, logger.GetLogger())
	downloadTokenBackgroundService = download_token.NewBackgroundService(downloadTokenService, logger.GetLogger())

	restoreStreamWriter = restore_stream.NewWriter(
		storages.GetStorageService(),
		util_encryption.GetFieldEncryptor(),
	)
}

func GetDownloadTokenService() *download_token.Service {
	return downloadTokenService
}

func GetRestoreTokenService() *restore_token.Service {
	return restoreTokenService
}

func GetDownloadTokenBackgroundService() *download_token.BackgroundService {
	return downloadTokenBackgroundService
}

func GetRestoreStreamWriter() *restore_stream.Writer {
	return restoreStreamWriter
}
