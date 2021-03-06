package funds

import (
	"encoding/json"
	"net/http"

	"github.com/weimpact/covid-ed/logger"
)

func Lister(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fs, err := svc.ListFunds(r.Context())
		if err != nil {
			logger.Errorf("[Funds] error listing funds: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(fs) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if err := json.NewEncoder(w).Encode(fs); err != nil {
			logger.Errorf("[Funds] error writing response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
