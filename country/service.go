package country

import (
	"context"

	"github.com/weimpact/covid-ed/pkg/client"
)

type Service struct {
	cli client.Client
}

type Filter struct {
	top    int
	deaths bool
}

func (s Service) GetCountriesCasesAggregated(ctx context.Context, i interval, countries []string) (client.CountriesCases, error) {
	var aggregatedCases client.CountriesCases
	switch i {
	case daily:
		return s.CountriesCasesAggregatedDaily(ctx, countries)
	case weekly:
		return s.CountriesCasesAggregatedWeekly(ctx, countries)
	}
	return aggregatedCases, nil
}

func NewService(cli client.Client) Service {
	return Service{cli: cli}
}

func (s Service) CountriesCasesAggregatedDaily(ctx context.Context, countries []string) (client.CountriesCases, error) {
	cases, err := s.cli.GetCountriesCasesByDay(ctx, countries)
	if err != nil {
		return client.CountriesCases{}, err
	}
	return cases, nil
}

func (s Service) CountriesCasesAggregatedWeekly(ctx context.Context, countries []string) (client.CountriesCases, error) {
	cases, err := s.cli.GetCountriesCasesByWeek(ctx, countries)
	if err != nil {
		return client.CountriesCases{}, err
	}
	return cases, nil
}

func (s Service) GetCountriesData(ctx context.Context, filter Filter) (client.CountriesData, error) {
	if filter.top > 0 {
		return s.cli.GetTopCountriesData(ctx, filter.top)
	}
	if filter.deaths {
		return s.cli.GetCountriesDataWithDeaths(ctx)
	}
	return client.CountriesData{}, nil
}

func (s Service) GetCountries(ctx context.Context) (client.Countries, error) {
	return s.cli.GetCountries(ctx)
}
