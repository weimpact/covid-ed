package store

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/text/language"
)

type Store struct {
	DB *sqlx.DB
}

type Myth struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Locale      string    `db:"locale"`
	CreatedAt   time.Time `db:"created_at"`
}

type Fact struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Locale      string    `db:"locale"`
	CreatedAt   time.Time `db:"created_at"`
}

type FactAndMyth struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Locale      string `db:"locale"`
	Myth        `db:"myth"`
	CreatedAt   time.Time `db:"created_at"`
}

func (s *Store) FetchFacts(ctx context.Context, locale language.Tag) ([]Fact, error) {
	var facts []Fact
	query := s.DB.Rebind("select * from facts where locale=?")
	err := s.DB.SelectContext(ctx, &facts, query, locale.String())
	return facts, err
}

func (s *Store) FetchFactsAndMyths(ctx context.Context, locale language.Tag) ([]FactAndMyth, error) {
	var factsMyths []FactAndMyth
	query := `
            select f.*,
            m.title "myth.title", m.description "myth.description", m.created_at "myth.created_at"
            From
            facts f join myths m on f.id=m.fact_id
            and f.locale = ?
            `
	query = s.DB.Rebind(query)
	err := s.DB.SelectContext(ctx, &factsMyths, query, locale.String())
	return factsMyths, err
}
