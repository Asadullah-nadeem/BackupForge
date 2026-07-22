package usecases_logical_postgresql

import (
	"BackupForge-backend/internal/features/encryption/secrets"
	"BackupForge-backend/internal/util/encryption"
	"BackupForge-backend/internal/util/logger"
)

var createPostgresqlBackupUsecase = &CreatePostgresqlBackupUsecase{
	logger.GetLogger(),
	secrets.GetSecretKeyService(),
	encryption.GetFieldEncryptor(),
}

func GetCreatePostgresqlBackupUsecase() *CreatePostgresqlBackupUsecase {
	return createPostgresqlBackupUsecase
}
