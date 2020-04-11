package store

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	DB *sqlx.DB
}

type Myth struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

type Fact struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

type FactAndMyth struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Myth        `db:"myth"`
	CreatedAt   time.Time `db:"created_at"`
}

func (s *Store) FetchFacts(ctx context.Context) ([]Fact, error) {
	var facts []Fact
	query := s.DB.Rebind("select * from facts")
	err := s.DB.SelectContext(ctx, &facts, query)
	return facts, err
}

func (s *Store) FetchFactsAndMyths(ctx context.Context) ([]FactAndMyth, error) {
	var factsMyths []FactAndMyth
	query := `
            select f.*,
            m.title "myth.title", m.description "myth.description", m.created_at "myth.created_at"
            From
            facts f join myths m on f.id=m.fact_id
            `
	err := s.DB.SelectContext(ctx, &factsMyths, query)
	return factsMyths, err
}
