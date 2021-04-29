package facts

import (
	"context"

	"github.com/weimpact/covid-ed/logger"
	"github.com/weimpact/covid-ed/store"
	"golang.org/x/text/language"
)

type Service struct {
	store store.Store
}

type Myth struct {
	ID          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Locale      string `json:"locale,omitempty"`
}

type Fact struct {
	ID          int       `json:"-"`
	Title       string    `json:"title"`
	Articles    []Article `json:"articles,omitempty"`
	Description string    `json:"description"`
	Locale      string    `json:"locale,omitempty"`
}

func (f *Fact) AddArticle(a Article) {
	f.Articles = append(f.Articles, a)
}

type Article struct {
	ID     int    `json:"-"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	FactID int    `json:"-"`
	Locale string `json:"locale,omitempty"`
}

type FactAndMyth struct {
	Fact `json:"fact"`
	Myth `json:"myth"`
}

func (s Service) getLocale(l string) language.Tag {
	defaultLang := language.AmericanEnglish
	if l == "" {
		return defaultLang
	}
	lang, err := language.Parse(l)
	if err != nil {
		logger.Infof("[Facts.Service] unsupported locale: %s err: %v; returning default en-US", l, err)
		return defaultLang
	}
	return lang
}

func (s Service) ListFacts(ctx context.Context, lang string) ([]Fact, error) {
	locale := s.getLocale(lang)
	fs, err := s.store.FetchFacts(ctx, locale)
	if err != nil || len(fs) == 0 {
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
		art := Article{a.ID, a.Title, a.URL, a.FactID, a.Locale}
		data[key].AddArticle(art)
	}
	return data, nil
}

func (s Service) ListFactWithMyth(ctx context.Context, lang string) ([]FactAndMyth, error) {
	locale := s.getLocale(lang)
	fms, err := s.store.FetchFactsAndMyths(ctx, locale)
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
		art := Article{a.ID, a.Title, a.URL, a.FactID, a.Locale}
		data[key].AddArticle(art)
	}
	return data, nil
}

func NewService(s store.Store) Service {
	return Service{store: s}
}
