// Package app contains the main application
package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ak1m1tsu/urler/internal/app/api"
	"github.com/ak1m1tsu/urler/internal/pkg/config"
	"github.com/ak1m1tsu/urler/internal/pkg/service"
	"golang.org/x/sync/errgroup"
)

// Application represents the Urler application
type Application struct {
	ctx  context.Context
	cfg  *config.Config
	errg *errgroup.Group
}

// New creates Urler application
func New(configPath string) *Application {
	var (
		app = new(Application)
		err error
	)

	app.cfg, err = config.Parse(configPath)
	app.exitOnError(err)

	return app
}

// Run starts the main process of the application
func (app *Application) Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)

	app.errg, app.ctx = errgroup.WithContext(ctx)

	svc := service.New(app.cfg)

	router := api.Router(svc)

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", app.cfg.HTTP.Port),
		Handler:      router,
		IdleTimeout:  app.cfg.HTTP.Timeout.Idle,
		ReadTimeout:  app.cfg.HTTP.Timeout.Read,
		WriteTimeout: app.cfg.HTTP.Timeout.Write,
	}

	app.errg.Go(func() error {
		return srv.ListenAndServe()
	})

	app.errg.Go(func() error {
		<-ctx.Done()
		cancel()
		return srv.Shutdown(context.Background())
	})

	err := app.errg.Wait()
	if err != nil {
		fmt.Println(err)
	}

	app.exitOnError(err)
}

// exitOnError calls exitOnError if err isn't nil
func (app *Application) exitOnError(err error) {
	if err != nil {
		if errors.Is(err, context.Canceled) || errors.Is(err, http.ErrServerClosed) {
			return
		}

		fmt.Println(err)
		os.Exit(1)
	}
}
