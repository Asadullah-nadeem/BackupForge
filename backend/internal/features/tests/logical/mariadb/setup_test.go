package mariadb_logical

import (
	"os"
	"testing"

	logicaltesting "BackupForge-backend/internal/features/tests/logical/shared"
)

func TestMain(m *testing.M) {
	teardown := logicaltesting.Setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
