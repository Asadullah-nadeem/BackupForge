package users_services

import (
	"BackupForge-backend/internal/features/email"
	"BackupForge-backend/internal/features/encryption/secrets"
	users_repositories "BackupForge-backend/internal/features/users/repositories"
)

var userService = &UserService{
	users_repositories.GetUserRepository(),
	secrets.GetSecretKeyService(),
	settingsService,
	nil,
	email.GetEmailSMTPSender(),
	users_repositories.GetPasswordResetRepository(),
}

var settingsService = &SettingsService{
	users_repositories.GetUsersSettingsRepository(),
	nil,
}

var managementService = &UserManagementService{
	users_repositories.GetUserRepository(),
	nil,
}

func GetUserService() *UserService {
	return userService
}

func GetSettingsService() *SettingsService {
	return settingsService
}

func GetManagementService() *UserManagementService {
	return managementService
}
