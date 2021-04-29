package main

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/weimpact/covid-ed/config"
	"github.com/weimpact/covid-ed/country"
	"github.com/weimpact/covid-ed/facts"
	"github.com/weimpact/covid-ed/funds"
	"github.com/weimpact/covid-ed/lang"
	"github.com/weimpact/covid-ed/media"
	"github.com/weimpact/covid-ed/pkg/client"
	"github.com/weimpact/covid-ed/store"

	gomw "github.com/devdinu/middlers"
	"github.com/gorilla/mux"
)

func server() (*mux.Router, error) {
	m := mux.NewRouter()
	m.Use(mux.MiddlewareFunc(contentWriter))
	m.Use(mux.CORSMethodMiddleware(m))
	m.Use(mux.MiddlewareFunc(accessController))

	cli, err := client.New()
	if err != nil {
		return nil, err
	}
	countryService := country.NewService(cli)
	db, err := NewDB(config.Database())
	if err != nil {
		return nil, err
	}
	store := store.Store{DB: db}

	factService := facts.NewService(store)
	fundService := funds.NewService(store)
	mediaService := media.NewService(store)

	m.HandleFunc("/ping", PingHandler())

	m.HandleFunc("/countries/cases", gomw.RequestLogger(country.CaseLister(countryService))).Methods(http.MethodGet, http.MethodOptions)
	m.HandleFunc("/countries/cases/aggregated", gomw.RequestLogger(country.CountriesAggregatedCasesHandler(countryService))).Methods(http.MethodGet, http.MethodOptions)
	m.HandleFunc("/facts", gomw.RequestLogger(facts.Lister(factService))).Methods(http.MethodGet, http.MethodOptions)
	m.HandleFunc("/facts_myths", gomw.RequestLogger(facts.ListWithFacts(factService))).Methods(http.MethodGet, http.MethodOptions)
	m.HandleFunc("/languages", gomw.RequestLogger(lang.ListSupportedLanguages())).Methods(http.MethodGet, http.MethodOptions)
	m.HandleFunc("/countries", gomw.RequestLogger(country.List(countryService))).Methods(http.MethodGet, http.MethodOptions)
	m.HandleFunc("/funds", gomw.RequestLogger(funds.Lister(fundService))).Methods(http.MethodGet, http.MethodOptions)
	m.HandleFunc("/media", gomw.RequestLogger(media.Lister(mediaService))).Methods(http.MethodGet, http.MethodOptions)

	return m, nil
}

func accessController(next http.Handler) http.Handler {
	domains := config.AccessControlAllowOrigin()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, allowed := range domains {
			if r.Header.Get("Origin") == allowed {
				w.Header().Set("Access-Control-Allow-Origin", allowed)
				w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Auth-Code, Set-Cookie")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
		}
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
		}
		next.ServeHTTP(w, r)
	})
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

func NewDB(cfg config.DB) (*sqlx.DB, error) {
	var err error
	db, err := sqlx.Open(cfg.Driver, cfg.URL())
	if err != nil {
		return nil, fmt.Errorf("error opening conn to db: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.MaxConnLifetime())
	return db, nil
}
