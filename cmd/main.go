package main

import (
	"context"
	"errors"
	"fmt"
	"mac-watcher/config"
	"mac-watcher/internal/domain"
	"mac-watcher/internal/repository"
	"mac-watcher/internal/service"
	"mac-watcher/pkg/database"
	"mac-watcher/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

const (
	fileName = ".env"
)

func main() {

	zapLogger, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}

	ctx, cancelContext := context.WithCancel(context.Background())

	conf, err := config.NewConfig(fileName)
	if err != nil {
		zapLogger.Error("error init config", zap.Error(err))
		return
	}

	newDatabase, err := database.NewDatabase(conf.DBConfig)
	if err != nil {
		zapLogger.Error("error init database", zap.Error(err))
		return
	}
	defer newDatabase.Close()

	if err := database.Migrate(newDatabase, zapLogger); !errors.Is(err, domain.ErrExistsTable) {
		zapLogger.Error("error migarate to database", zap.Error(err))
	}

	repositories := repository.NewRepositories(newDatabase)
	service.NewServices(ctx, conf, zapLogger, repositories)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	cancelContext()

	for i := 5; i > 0; i-- {
		time.Sleep(time.Second)
		fmt.Println(i)
	}

	zapLogger.Info("application closed")
}
