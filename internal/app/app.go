// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/fire9900/go-forum/config"
	"github.com/fire9900/go-forum/internal/controller/http"
	"github.com/fire9900/go-forum/internal/repo/persistent"
	"github.com/fire9900/go-forum/internal/usecase/thread"
	"github.com/fire9900/go-forum/pkg/httpserver"
	"github.com/fire9900/go-forum/pkg/logger"
	"github.com/fire9900/go-forum/pkg/postgres"
	"os"
	"os/signal"
	"syscall"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) { //nolint: gocyclo,cyclop,funlen,gocritic,nolintlint
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use-Case
	threadUseCase := thread.New(
		persistent.NewThreadRepo(pg),
	)

	// gRPC Server
	//grpcServer := grpcserver.New(l, grpcserver.Port(cfg.GRPC.Port))
	//grpc.NewRouter(grpcServer.App, translationUseCase, l)

	// HTTP Server
	httpServer := httpserver.New(l, httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))
	http.NewRouter(httpServer.App, cfg, threadUseCase, l)

	// Start servers
	//grpcServer.Start()
	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: %s", s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
		//case err = <-grpcServer.Notify():
		//	l.Error(fmt.Errorf("app - Run - grpcServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	//err = grpcServer.Shutdown()
	//if err != nil {
	//	l.Error(fmt.Errorf("app - Run - grpcServer.Shutdown: %w", err))
	//}
}
