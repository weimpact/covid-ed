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
		baseURL: "https://api.covid19api.com/",
	}
	return cli, nil
}

type CountryStatus struct {
	CountryCode string `json:"country_code"`
	Country     string
	Lat         float32
	Lon         float32
	Cases       int
	Status      string
	Date        time.Time
}

type CaseStatus string

const (
	confirmed CaseStatus = "confirmed"
)

type Global struct{}

type Summary struct {
	Global
	Countries
}

func (c Client) Summary(ctx context.Context) (Summary, error) {
	url := fmt.Sprintf("%s/summary", c.baseURL)
	resp, err := c.httpCli.Get(url)
	if err != nil {
		return Summary{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return Summary{}, fmt.Errorf("error getting response code: %d", resp.StatusCode)
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
		return nil, fmt.Errorf("error getting response code: %d", resp.StatusCode)
	}
	var cases CasesByDay
	if err := json.NewDecoder(resp.Body).Decode(&cases); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return cases, err
}
