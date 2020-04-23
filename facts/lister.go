package facts

import (
	"encoding/json"
	"net/http"

	"github.com/weimpact/covid-ed/logger"
)

func Lister(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		locale := r.URL.Query().Get("locale")

		fs, err := svc.ListFacts(r.Context(), locale)
		if err != nil {
			logger.Errorf("[Facts] error listing facts: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(fs) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if err := json.NewEncoder(w).Encode(fs); err != nil {
			logger.Errorf("[Facts] error writing response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func ListWithFacts(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		locale := r.URL.Query().Get("locale")

		fms, err := svc.ListFactWithMyth(r.Context(), locale)
		if err != nil {
			logger.Errorf("[FactsMyths] error listing facts: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(fms) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if err := json.NewEncoder(w).Encode(fms); err != nil {
			logger.Errorf("[FactMyths] error writing response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
