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

func New(configPath string) *Application {
	var (
		app = new(Application)
		err error
	)

	app.cfg, err = config.Parse(configPath)
	app.exitOnError(err)

	return app
}

func (app *Application) Run() {
	fmt.Println("do nothing")
}

func (app *Application) exitOnError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
