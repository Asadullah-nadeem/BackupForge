package usecases_mysql

import (
	"BackupForge-backend/internal/features/encryption/secrets"
	"BackupForge-backend/internal/util/logger"
)

var restoreMysqlBackupUsecase = &RestoreMysqlBackupUsecase{
	logger.GetLogger(),
	secrets.GetSecretKeyService(),
}

func GetRestoreMysqlBackupUsecase() *RestoreMysqlBackupUsecase {
	return restoreMysqlBackupUsecase
}
