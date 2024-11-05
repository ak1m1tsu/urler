// Package app contains the main application
package app

import (
	"fmt"
	"os"

	"github.com/ak1m1tsu/urler/internal/pkg/config"
)

// Application represents the Urler application
type Application struct {
	cfg *config.Config
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
	fmt.Println("do nothing")
}

// exitOnError calls exitOnError if err isn't nil
func (app *Application) exitOnError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
