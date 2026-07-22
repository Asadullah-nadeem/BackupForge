package usecases_logical_mysql

import (
	"BackupForge-backend/internal/features/encryption/secrets"
	"BackupForge-backend/internal/util/encryption"
	"BackupForge-backend/internal/util/logger"
)

var createMysqlBackupUsecase = &CreateMysqlBackupUsecase{
	logger.GetLogger(),
	secrets.GetSecretKeyService(),
	encryption.GetFieldEncryptor(),
}

func GetCreateMysqlBackupUsecase() *CreateMysqlBackupUsecase {
	return createMysqlBackupUsecase
}
