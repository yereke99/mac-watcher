package service

import (
	"context"
	"mac-watcher/config"
	"mac-watcher/internal/domain"
	"mac-watcher/internal/repository"

	"go.uber.org/zap"
)

type IMacWatcherService interface {
	// TODO: ?
	listenMacReboot(clouds []*domain.Cloud)
	// TODO: ?
	listenMacRecovery(clouds []*domain.Cloud)
	// TODO: ?
	sendMessage(message string) error
	// TODO: ?
	macPing(cloud *domain.Cloud) error
}

type Services struct {
	MacWatcherService IMacWatcherService
}

func NewServices(ctx context.Context, appConfig *config.Config, zapLogger *zap.Logger, repositories *repository.Repositories) *Services {

	services := &Services{
		MacWatcherService: NewMacWatcherService(ctx, appConfig, zapLogger, repositories),
	}
	return services
}
