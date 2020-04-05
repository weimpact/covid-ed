package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

var app Application

func MustLoad() {
	var errs []error

	if err := envconfig.Process("APP", &app.server); err != nil {
		errs = append(errs, err)
	}

	if len(errs) != 0 {
		log.Fatalf("Error loading configuration: %v", errs)
	}
}
