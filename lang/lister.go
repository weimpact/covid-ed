package lang

import (
	"encoding/json"
	"net/http"

	"github.com/weimpact/covid-ed/logger"
	"golang.org/x/text/language"
)

type Lang struct {
	Tag       language.Tag `json:"tag"`
	Languague string       `json:"language"`
}
type Languages []Lang

func ListSupportedLanguages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		langs := []Lang{
			{language.AmericanEnglish, "English"},
			{language.Tamil, "Tamil"},
			{language.Hindi, "Hindi"},
		}
		if err := json.NewEncoder(w).Encode(langs); err != nil {
			logger.Errorf("[Lang] error writing response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
