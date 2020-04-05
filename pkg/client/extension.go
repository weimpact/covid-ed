package client

import (
	"context"
	"sort"
)

type TopCountries struct {
	Total     int
	Countries []Country `json:"Countries"`
}

type byTotalConfirmed []Country

func (c byTotalConfirmed) Len() int           { return len(c) }
func (c byTotalConfirmed) Less(i, j int) bool { return c[i].TotalConfirmed > c[j].TotalConfirmed }
func (c byTotalConfirmed) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func (c Client) GetTopCountriesData(ctx context.Context, top int) (TopCountries, error) {
	summary, err := c.Summary(ctx)
	if err != nil {
		return TopCountries{}, err
	}
	maxConfirmed := byTotalConfirmed(summary.Countries)
	sort.Sort(maxConfirmed)
	return TopCountries{Total: top, Countries: maxConfirmed[:top]}, nil
}
