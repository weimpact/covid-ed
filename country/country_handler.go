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
	GetTopCountriesData(context.Context, int) (client.TopCountries, error)
}

func TopN(svc service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		top := 5
		topQ := r.URL.Query().Get("top")
		if topQ != "" {
			var err error
			top, err = strconv.Atoi(topQ)
			if err != nil {
				logger.Infof("[TopN] Invalid top query param request: %v", err)
				return
			}
		}

		resp, err := svc.GetTopCountriesData(r.Context(), top)
		if err != nil {
			logger.Errorf("[TopN] error fetching countries data: %v", err)
			return
		}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			logger.Errorf("[TopN] error writing response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
