package usecases_mongodb

import (
	encryption_secrets "BackupForge-backend/internal/features/encryption/secrets"
	"BackupForge-backend/internal/util/logger"
)

var restoreMongodbBackupUsecase = &RestoreMongodbBackupUsecase{
	logger.GetLogger(),
	encryption_secrets.GetSecretKeyService(),
}

func GetRestoreMongodbBackupUsecase() *RestoreMongodbBackupUsecase {
	return restoreMongodbBackupUsecase
}
