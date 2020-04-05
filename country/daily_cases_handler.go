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
	GetCountriesCasesByDay(context.Context, []string) (client.CountriesCases, error)
}

func CountriesAggregatedCasesHandler(svc aggregatedService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		countriesQ := r.URL.Query().Get("countries")
		cns := strings.Split(countriesQ, ",")
		if len(cns) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp, err := svc.GetCountriesCasesByDay(r.Context(), cns)
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
