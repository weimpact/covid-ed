package media

import (
	"context"

	"github.com/weimpact/covid-ed/store"
)

type Service struct {
	store store.Store
}

func (s Service) FetchMedia(ctx context.Context, t MediaType, category string) ([]Media, error) {
	mediaFiles, err := s.store.FetchMedia(ctx, string(t), category)
	if err != nil {
		return nil, err
	}

	medias := make([]Media, len(mediaFiles))
	for i, m := range mediaFiles {
		media := Media{
			URL:         m.URL,
			Kind:        m.Kind,
			Category:    m.Category,
			Title:       m.Title,
			Description: m.Description,
		}
		medias[i] = media
	}
	return medias, nil
}

func NewService(st store.Store) Service {
	return Service{store: st}
}
