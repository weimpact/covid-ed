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

type Countries []Country

type summary struct {
	NewConfirmed   int64
	TotalConfirmed int64
	NewDeaths      int64
	TotalDeaths    int64
	NewRecovered   int64
	TotalRecovered int64
}

type Summary struct {
	Global summary
	Countries
}
