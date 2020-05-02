package funds

import (
	"context"
	"fmt"

	"github.com/weimpact/covid-ed/config"
	"github.com/weimpact/covid-ed/store"
)

type Service struct {
	store store.Store
}

type Fund struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Website     string `json:"website"`
	ImageURL    string `json:"image_url"`
	DonateURL   string `json:"donate_url"`
}

func (s Service) ListFunds(ctx context.Context) ([]Fund, error) {
	fs, err := s.store.FetchFunds(ctx)
	if err != nil {
		return nil, err
	}
	allFunds := make([]Fund, len(fs))
	for i, f := range fs {
		f := Fund{Title: f.Title, Description: f.Description, Website: f.Website,
			ImageURL:  s.buildStaticURL(f.ImageURL),
			DonateURL: f.DonateURL,
		}
		allFunds[i] = f
	}
	return allFunds, nil
}

func (s Service) buildStaticURL(name string) string {
	domain := config.StaticServerDomain()
	return fmt.Sprintf("%s/%s", domain, name)
}

func NewService(s store.Store) Service {
	return Service{store: s}
}
