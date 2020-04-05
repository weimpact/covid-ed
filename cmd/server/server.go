package main

import (
	"net/http"

	"github.com/weimpact/covid-ed/country"
	"github.com/weimpact/covid-ed/pkg/client"

	gomw "github.com/devdinu/middlers"
	"github.com/gorilla/mux"
)

func server() (*mux.Router, error) {
	m := mux.NewRouter()
	m.Use(mux.MiddlewareFunc(contentWriter))

	cli, err := client.New()
	if err != nil {
		return nil, err
	}

	m.HandleFunc("/ping", PingHandler())

	m.HandleFunc("/countries/cases", gomw.RequestLogger(country.TopN(cli)))
	m.HandleFunc("/countries/cases/aggregated", gomw.RequestLogger(country.CountriesAggregatedCasesHandler(cli)))

	return m, nil
}

func contentWriter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"response":"pong"}`))
	}
}
