package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

var app Application

func MustLoad() {
	var errs []error

	if err := envconfig.Process("", &app.server); err != nil {
		errs = append(errs, err)
	}

	if err := envconfig.Process("DB", &app.db); err != nil {
		errs = append(errs, err)
	}

	if err := envconfig.Process("STATIC_SERVER", &app.StaticServer); err != nil {
		errs = append(errs, err)
	}

	if len(errs) != 0 {
		log.Fatalf("Error loading configuration: %v", errs)
	}
	fmt.Printf("%+v", app)
}
