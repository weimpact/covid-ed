package main

import (
	"log"
	"net/http"

	"github.com/weimpact/covid-ed/config"
	"github.com/weimpact/covid-ed/logger"
)

func main() {
	config.MustLoad()
	router, err := server()
	if err != nil {
		logger.Fatalf("error creating server: %v", err)
	}

	addr := config.AppAddress()
	logger.Infof("listening on address %s", addr)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatalf("error listening for rerquests on port: %s err: %v", addr, err)
	}
}
