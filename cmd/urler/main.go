// Package main contains urler entrypoint
package main

import (
	"flag"

	"github.com/ak1m1tsu/urler/internal/app"
)

var path string

func init() {
	flag.StringVar(&path, "config", "config.yaml", "path to config file for urler")
	flag.Parse()
}

func main() {
	app.New(path).Run()
}
