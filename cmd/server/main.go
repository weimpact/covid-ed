package main

import (
	"log"
	"net/http"
	"weimpact/covid-ed/config"

	"github.com/gojek/kat/logger"
)

func main() {
	config.MustLoad()

	router, err := server()
	if err != nil {
		logger.Fatalf("error creating server: %v", err)
	}
	err = http.ListenAndServe(config.AppAddress(), router)
	if err != nil {
		log.Fatalf("error listening for rerquests on port: %s err: %v", config.AppAddress(), err)
	}
}
