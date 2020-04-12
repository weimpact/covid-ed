package facts

import (
	"context"

	"github.com/weimpact/covid-ed/store"
)

type Service struct {
	store store.Store
}

type Myth struct {
	ID          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Fact struct {
	ID          int       `json:"-"`
	Title       string    `json:"title"`
	Articles    []Article `json:"articles,omitempty"`
	Description string    `json:"description"`
}

func (f *Fact) AddArticle(a Article) {
	f.Articles = append(f.Articles, a)
}

type Article struct {
	ID     int    `json:"-"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	FactID int    `json:"-"`
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
	data := make([]Fact, len(fs))
	var factIDs []int
	dataMap := make(map[int]int)
	for i, f := range fs {
		d := Fact{ID: f.ID, Title: f.Title, Description: f.Description}
		factIDs = append(factIDs, f.ID)
		dataMap[f.ID] = i
		data[i] = d
	}
	articles, err := s.store.FetchArticles(ctx, factIDs)
	if err != nil {
		return nil, err
	}
	for _, a := range articles {
		key := dataMap[a.FactID]
		art := Article{a.ID, a.Title, a.URL, a.FactID}
		data[key].AddArticle(art)
	}
	return data, nil
}

func (s Service) ListFactWithMyth(ctx context.Context) ([]FactAndMyth, error) {
	fms, err := s.store.FetchFactsAndMyths(ctx)
	if err != nil {
		return nil, err
	}
	data := make([]FactAndMyth, len(fms))
	var factIDs []int
	dataMap := make(map[int]int)
	for i, fm := range fms {
		fact := Fact{ID: fm.ID, Title: fm.Title, Description: fm.Description}
		myth := Myth{Title: fm.Myth.Title, Description: fm.Myth.Description}
		factIDs = append(factIDs, fm.ID)
		dataMap[fm.ID] = i
		data[i] = FactAndMyth{Fact: fact, Myth: myth}
	}
	articles, err := s.store.FetchArticles(ctx, factIDs)
	if err != nil {
		return nil, err
	}
	for _, a := range articles {
		key := dataMap[a.FactID]
		art := Article{a.ID, a.Title, a.URL, a.FactID}
		data[key].AddArticle(art)
	}
	return data, nil
}

func NewService(s store.Store) Service {
	return Service{store: s}
}
