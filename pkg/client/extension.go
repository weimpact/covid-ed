package client

import (
	"context"
	"fmt"
	"sort"
	"time"
)

type CountriesData struct {
	Global    summary
	Countries []Country
}

type byTotalConfirmed []Country

func (c byTotalConfirmed) Len() int           { return len(c) }
func (c byTotalConfirmed) Less(i, j int) bool { return c[i].TotalConfirmed > c[j].TotalConfirmed }
func (c byTotalConfirmed) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

type byTotalDeaths []Country

func (c byTotalDeaths) Len() int           { return len(c) }
func (c byTotalDeaths) Less(i, j int) bool { return c[i].TotalDeaths > c[j].TotalDeaths }
func (c byTotalDeaths) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func (c Client) GetTopCountriesData(ctx context.Context, top int) (CountriesData, error) {
	summary, err := c.Summary(ctx)
	if err != nil {
		return CountriesData{}, err
	}
	maxConfirmed := byTotalConfirmed(summary.Countries)
	sort.Sort(maxConfirmed)
	return CountriesData{Countries: maxConfirmed[:top]}, nil
}

func (c Client) GetCountriesDataWithDeaths(ctx context.Context) (CountriesData, error) {
	resp, err := c.Summary(ctx)
	if err != nil {
		return CountriesData{}, err
	}
	var global summary
	var data []Country
	for _, c := range resp.Countries {
		if c.TotalDeaths > 0 {
			data = append(data, c)
			global.TotalDeaths += c.TotalDeaths
			global.TotalConfirmed += c.TotalConfirmed
			global.TotalRecovered += c.TotalRecovered
		}
	}
	countriesWithDeaths := byTotalDeaths(data)
	sort.Sort(countriesWithDeaths)
	return CountriesData{Global: global, Countries: countriesWithDeaths}, nil
}

type CountriesCases struct {
	CountriesCases []AggregateCountryStatus `json:"Countries"`
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
