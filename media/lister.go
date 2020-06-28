package media

import (
	"encoding/json"
	"net/http"

	"github.com/weimpact/covid-ed/logger"
)

type MediaType string
type Category string

const (
	Image   MediaType = "image"
	Video             = "video"
	Article           = "article"
	All               = "all"
)

type Media struct {
	Title       string `json:"title",omitempty`
	Description string `json:"description",omitempty`
	URL         string `json:"url"`
	Kind        string `json:"kind,omitempty"`
	Category    string `json:"category,omitempty"`
}

type MediaResponse struct {
	Data []Media `json:"media,omitempty"`
}

func Lister(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mtq := r.URL.Query().Get("type")
		mt := parseMediaType(mtq)
		medias, err := svc.FetchMedia(r.Context(), mt, "all")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Errorf("[Media] error listing media: %v", err)
		}
		resp := MediaResponse{Data: medias}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func parseMediaType(mt string) MediaType {
	switch mt {
	case "image":
		return Image
	case "video":
		return Video
	default:
		return All
	}
}
