package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	httpCli http.Client
	baseURL string
}

func New() (Client, error) {
	cli := Client{
		httpCli: http.Client{},
		baseURL: "https://api.covid19api.com",
	}
	return cli, nil
}

type CountryStatus struct {
	CountryCode string
	Country     string
	Lat         string
	Lon         string
	Cases       int
	Status      string
	Date        time.Time
}

type CaseStatus string

const (
	confirmed CaseStatus = "confirmed"
)

func (c Client) Summary(ctx context.Context) (Summary, error) {
	url := fmt.Sprintf("%s/summary", c.baseURL)
	resp, err := c.httpCli.Get(url)
	if err != nil {
		return Summary{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return Summary{}, fmt.Errorf("summary: error getting response code: %d, url:%s", resp.StatusCode, url)
	}
	var summary Summary
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		return Summary{}, err
	}
	defer resp.Body.Close()
	return summary, err
}

type CasesByDay []CountryStatus

func (c Client) CountryLiveCasesEveryday(ctx context.Context, country string, status CaseStatus) (CasesByDay, error) {
	url := fmt.Sprintf("%s/country/%s/status/%s/live", c.baseURL, country, status)

	resp, err := c.httpCli.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("country_cases: error getting response code: %d url:%s", resp.StatusCode, url)
	}
	var cases CasesByDay
	if err := json.NewDecoder(resp.Body).Decode(&cases); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return cases, err
}

func (c Client) DayOneCountryLiveCasesEveryday(ctx context.Context, country string, status CaseStatus) (CasesByDay, error) {
	url := fmt.Sprintf("%s/dayone/country/%s/status/%s/live", c.baseURL, country, status)

	resp, err := c.httpCli.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("country dayone: error getting response code: %d", resp.StatusCode)
	}
	var cases CasesByDay
	if err := json.NewDecoder(resp.Body).Decode(&cases); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return cases, err
}

func (c Client) GetCountries(ctx context.Context) (Countries, error) {
	url := fmt.Sprintf("%s/countries", c.baseURL)
	resp, err := c.httpCli.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("countries: error getting response code: %d url:%s", resp.StatusCode, url)
	}
	var countries Countries
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return countries, err
}
