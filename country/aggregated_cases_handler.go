package country

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/weimpact/covid-ed/logger"
	"github.com/weimpact/covid-ed/pkg/client"
)

type aggregatedService interface {
	GetCountriesCasesAggregated(context.Context, interval, []string) (client.CountriesCases, error)
}
type interval string

const (
	daily  interval = "daily"
	weekly interval = "weekly"
)

func CountriesAggregatedCasesHandler(svc aggregatedService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		countriesQ := r.URL.Query().Get("countries")
		cns := strings.Split(countriesQ, ",")
		if countriesQ == "" {
			logger.Errorf("[CountriesAggregated] No countries information's available")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		intervalQ := r.URL.Query().Get("interval")

		resp, err := svc.GetCountriesCasesAggregated(r.Context(), getInterval(intervalQ), cns)
		if err != nil {
			logger.Errorf("[CountriesAggregated] couldn't get status by day for cns: %s error: %v", cns, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			logger.Errorf("[CountriesAggregated] error writing response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func getInterval(i string) interval {
	switch i {
	case "daily":
		return daily
	case "weekly":
		return weekly
	default:
		return weekly
	}
}
