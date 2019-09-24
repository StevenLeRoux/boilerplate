package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/stevenleroux/boilerplate/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.WithError(err).Fatal("Could not execute boilerplate")
	}
}
