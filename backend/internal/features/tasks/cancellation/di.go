package task_cancellation

import (
	"context"
	"sync"

	"github.com/google/uuid"

	cache_utils "BackupForge-backend/internal/util/cache"
	"BackupForge-backend/internal/util/logger"
)

var taskCancelManager = &TaskCancelManager{
	sync.RWMutex{},
	make(map[uuid.UUID]context.CancelFunc),
	cache_utils.NewPubSubManager(),
	logger.GetLogger(),
}

func GetTaskCancelManager() *TaskCancelManager {
	return taskCancelManager
}

var SetupDependencies = sync.OnceFunc(func() {
	taskCancelManager.StartSubscription()
})
