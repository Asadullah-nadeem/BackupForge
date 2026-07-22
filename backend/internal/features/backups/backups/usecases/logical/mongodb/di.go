package usecases_logical_mongodb

import (
	encryption_secrets "BackupForge-backend/internal/features/encryption/secrets"
	"BackupForge-backend/internal/util/encryption"
	"BackupForge-backend/internal/util/logger"
)

var createMongodbBackupUsecase = &CreateMongodbBackupUsecase{
	logger.GetLogger(),
	encryption_secrets.GetSecretKeyService(),
	encryption.GetFieldEncryptor(),
}

func GetCreateMongodbBackupUsecase() *CreateMongodbBackupUsecase {
	return createMongodbBackupUsecase
}
