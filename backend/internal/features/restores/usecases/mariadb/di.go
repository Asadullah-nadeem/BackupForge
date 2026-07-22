package usecases_mariadb

import (
	"BackupForge-backend/internal/features/encryption/secrets"
	"BackupForge-backend/internal/util/logger"
)

var restoreMariadbBackupUsecase = &RestoreMariadbBackupUsecase{
	logger.GetLogger(),
	secrets.GetSecretKeyService(),
}

func GetRestoreMariadbBackupUsecase() *RestoreMariadbBackupUsecase {
	return restoreMariadbBackupUsecase
}
