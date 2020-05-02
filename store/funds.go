package store

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type Fund struct {
	ID          string    `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Website     string    `db:"website"`
	ImageURL    string    `db:"image_url"`
	DonateURL   string    `db:"donate_url"`
	CreatedAt   time.Time `db:"created_at"`
}

func (s *Store) FetchFunds(ctx context.Context) ([]Fund, error) {
	var funds []Fund
	query, args, err := sqlx.In("select * from funds")
	if err != nil {
		return nil, err
	}
	query = s.DB.Rebind(query)
	if err = s.DB.SelectContext(ctx, &funds, query, args...); err != nil {
		return nil, err
	}
	return funds, nil
}
