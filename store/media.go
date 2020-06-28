package store

import (
	"context"
	"time"
)

type Media struct {
	ID          string    `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Category    string    `db:"category"`
	Kind        string    `db:"kind"`
	URL         string    `db:"url"`
	CreatedAt   time.Time `db:"created_at"`
}

func (s *Store) FetchMedia(ctx context.Context, kind, category string) ([]Media, error) {
	if category == "all" {
		return s.fetchAllMedia(ctx, kind)
	}
	var medias []Media
	query := s.DB.Rebind("select * from medias where category=? and kind=?")
	err := s.DB.SelectContext(ctx, &medias, query, category, kind)
	return medias, err
}

func (s *Store) fetchAllMedia(ctx context.Context, kind string) ([]Media, error) {
	var medias []Media
	query := s.DB.Rebind("select * from medias where kind=?")
	err := s.DB.SelectContext(ctx, &medias, query, kind)
	return medias, err
}
