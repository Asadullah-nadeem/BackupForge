package cloudflare_turnstile

import (
	"BackupForge-backend/internal/config"
)

var cloudflareTurnstileService = &CloudflareTurnstileService{
	config.GetEnv().CloudflareTurnstileSecretKey,
	config.GetEnv().CloudflareTurnstileSiteKey,
}

func GetCloudflareTurnstileService() *CloudflareTurnstileService {
	return cloudflareTurnstileService
}
