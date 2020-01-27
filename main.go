package main

import (
	"flag"
	"fmt"
	"os"
)

var app App

type App struct {
	listener string
	config   string
}

func init() {
	flag.StringVar(&app.listener, "a", "", "ip/port to listen on")
	flag.StringVar(&app.config, "c", "hookr.yml", "path to config file")
}

func main() {
	flag.Parse()
	s, err := NewServer(app.listener, app.config)
	exitOnErr(err...)
	s.run()
}

func exitOnErr(errs ...error) {
	errNotNil := false
	for _, err := range errs {
		if err == nil {
			continue
		}
		errNotNil = true
		fmt.Fprintf(os.Stderr, "ERROR: %s", err.Error())
	}
	if errNotNil {
		fmt.Print("\n")
		os.Exit(-1)
	}
}
