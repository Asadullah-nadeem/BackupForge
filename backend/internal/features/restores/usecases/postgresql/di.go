package usecases_postgresql

import (
	"BackupForge-backend/internal/features/encryption/secrets"
	"BackupForge-backend/internal/util/logger"
)

var restorePostgresqlBackupUsecase = &RestorePostgresqlBackupUsecase{
	logger.GetLogger(),
	secrets.GetSecretKeyService(),
}

func GetRestorePostgresqlBackupUsecase() *RestorePostgresqlBackupUsecase {
	return restorePostgresqlBackupUsecase
}
