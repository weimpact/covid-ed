package country

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/weimpact/covid-ed/logger"
	"github.com/weimpact/covid-ed/pkg/client"
)

type service interface {
	GetCountriesData(context.Context, Filter) (client.CountriesData, error)
}

func Lister(svc service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var filter Filter
		topQ := r.URL.Query().Get("top")
		if topQ != "" {
			var err error
			filter.top, err = strconv.Atoi(topQ)
			if err != nil {
				logger.Infof("[CountriesData] Invalid top query param request: %v", err)
				return
			}
		}

		deathsQ := r.URL.Query().Get("deaths")
		if deathsQ != "" && deathsQ == "true" {
			filter.deaths = true
		}

		resp, err := svc.GetCountriesData(r.Context(), filter)
		if err != nil {
			logger.Errorf("[CountriesData] error fetching countries data: %v", err)
			return
		}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			logger.Errorf("[CountriesData] error writing response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
