package client

type Country struct {
	Country        string
	CountryCode    string
	Slug           string
	NewConfirmed   int64
	TotalConfirmed int64
	NewDeaths      int64
	TotalDeaths    int64
	NewRecovered   int64
	TotalRecovered int64
	Date           string
}

type summary struct {
	NewConfirmed   int64
	TotalConfirmed int64
	NewDeaths      int64
	TotalDeaths    int64
	NewRecovered   int64
	TotalRecovered int64
}

type Summary struct {
	Global    summary
	Countries []Country
}

type Countries []CountryInfo
type CountryInfo struct {
	Country string
	Slug    string
	ISO2    string
}
