package store

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	DB *sqlx.DB
}
type Fact struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

func (s *Store) FetchFacts(ctx context.Context) ([]Fact, error) {
	var facts []Fact
	query := s.DB.Rebind("select * from facts")
	err := s.DB.SelectContext(ctx, &facts, query)
	return facts, err
}
