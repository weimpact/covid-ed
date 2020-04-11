package facts

import (
	"context"
	"time"

	"github.com/weimpact/covid-ed/store"
)

type Service struct {
	store store.Store
}

type Fact struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (s Service) ListFacts(ctx context.Context) ([]Fact, error) {
	fs, err := s.store.FetchFacts(ctx)
	if err != nil {
		return nil, err
	}
	var data []Fact
	for _, f := range fs {
		d := Fact{ID: f.ID, Description: f.Description, CreatedAt: f.CreatedAt}
		data = append(data, d)
	}
	return data, nil
}

func NewService(s store.Store) Service {
	return Service{store: s}
}
