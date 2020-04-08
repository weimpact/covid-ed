package client

import (
	"context"
	"fmt"
	"sort"
	"time"
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

type CountriesCases struct {
	CountriesCases []AggregateCountryStatus `Countries`
}

type AggregateCountryStatus struct {
	Country     string
	CountryCode string
	Cases       []int64
	Interval    string
	StartDate   time.Time
}

func (c Client) GetCountriesCasesByDay(ctx context.Context, countries []string) (CountriesCases, error) {
	var countriesCases []AggregateCountryStatus
	for _, cn := range countries {
		dailyCases, err := c.DayOneCountryLiveCasesEveryday(ctx, cn, confirmed)
		if err != nil {
			return CountriesCases{}, fmt.Errorf("error getting case for country :%s err: %v", cn, err)
		}
		aggregated := c.aggregateCasesByDay(ctx, dailyCases)
		countriesCases = append(countriesCases, aggregated)
	}
	return CountriesCases{CountriesCases: countriesCases}, nil
}

func (c Client) GetCountriesCasesByWeek(ctx context.Context, countries []string) (CountriesCases, error) {
	var countriesCases []AggregateCountryStatus
	for _, cn := range countries {
		dailyCases, err := c.DayOneCountryLiveCasesEveryday(ctx, cn, confirmed)
		if err != nil {
			return CountriesCases{}, fmt.Errorf("error getting case for country :%s err: %v", cn, err)
		}
		aggregated := c.aggregateCasesByWeek(ctx, dailyCases)
		countriesCases = append(countriesCases, aggregated)
	}
	return CountriesCases{CountriesCases: countriesCases}, nil
}

func (c Client) aggregateCasesByDay(ctx context.Context, dailyCases CasesByDay) AggregateCountryStatus {
	var aggregated AggregateCountryStatus
	if len(dailyCases) > 0 {
		aggregated.CountryCode = dailyCases[0].CountryCode
		aggregated.Country = dailyCases[0].Country
		aggregated.StartDate = dailyCases[0].Date
		aggregated.Interval = "Daily"
	}
	for _, day := range dailyCases {
		aggregated.Cases = append(aggregated.Cases, int64(day.Cases))
	}
	return aggregated
}

func (c Client) aggregateCasesByWeek(ctx context.Context, dailyCases CasesByDay) AggregateCountryStatus {
	var aggregated AggregateCountryStatus
	if len(dailyCases) > 0 {
		aggregated.CountryCode = dailyCases[0].CountryCode
		aggregated.Country = dailyCases[0].Country
		aggregated.StartDate = dailyCases[0].Date
		aggregated.Interval = "Weekly"
	}
	for i, day := range dailyCases {
		// Since the data returned have current data, we shouldn't sum it up
		// weeklyCases += day.Cases
		if (i+1)%7 == 0 {
			aggregated.Cases = append(aggregated.Cases, int64(day.Cases))
		}
	}
	l := len(dailyCases)
	if l%7 != 0 {
		aggregated.Cases = append(aggregated.Cases, int64(dailyCases[l-1].Cases))
	}
	return aggregated
}
