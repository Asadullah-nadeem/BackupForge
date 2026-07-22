package usecases_logical_mariadb

import (
	"BackupForge-backend/internal/features/encryption/secrets"
	"BackupForge-backend/internal/util/encryption"
	"BackupForge-backend/internal/util/logger"
)

var createMariadbBackupUsecase = &CreateMariadbBackupUsecase{
	logger.GetLogger(),
	secrets.GetSecretKeyService(),
	encryption.GetFieldEncryptor(),
}

func GetCreateMariadbBackupUsecase() *CreateMariadbBackupUsecase {
	return createMariadbBackupUsecase
}
