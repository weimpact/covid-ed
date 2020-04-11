package facts

import (
	"context"

	"github.com/weimpact/covid-ed/store"
)

type Service struct {
	store store.Store
}

type Myth struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Fact struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type FactAndMyth struct {
	Fact `json:"fact"`
	Myth `json:"myth"`
}

func (s Service) ListFacts(ctx context.Context) ([]Fact, error) {
	fs, err := s.store.FetchFacts(ctx)
	if err != nil {
		return nil, err
	}
	var data []Fact
	for _, f := range fs {
		d := Fact{ID: f.ID, Title: f.Title, Description: f.Description}
		data = append(data, d)
	}
	return data, nil
}

func (s Service) ListFactWithMyth(ctx context.Context) ([]FactAndMyth, error) {
	fms, err := s.store.FetchFactsAndMyths(ctx)
	if err != nil {
		return nil, err
	}
	var data []FactAndMyth
	for _, fm := range fms {
		fact := Fact{ID: fm.ID, Title: fm.Title, Description: fm.Description}
		myth := Myth{Title: fm.Myth.Title, Description: fm.Myth.Description}
		data = append(data, FactAndMyth{Fact: fact, Myth: myth})
	}
	return data, nil
}

func NewService(s store.Store) Service {
	return Service{store: s}
}
