package store

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type Article struct {
	ID        int       `db:"id"`
	Title     string    `db:"title"`
	FactID    int       `db:"fact_id"`
	URL       string    `db:"url"`
	Locale    string    `db:"locale"`
	CreatedAt time.Time `db:"created_at"`
}

func (s *Store) FetchArticles(ctx context.Context, factIDs []int) ([]Article, error) {
	var articles []Article
	query, args, err := sqlx.In("select * from articles where fact_id in (?)", factIDs)
	if err != nil {
		return nil, err
	}
	query = s.DB.Rebind(query)
	if err = s.DB.SelectContext(ctx, &articles, query, args...); err != nil {
		return nil, err
	}
	return articles, nil
}
