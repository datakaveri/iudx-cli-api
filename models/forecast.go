package models

type ForecastRequestBody struct {
	ForecastStart string `json:"forecastStart"`
	ForecastEnd   string `json:"forecastEnd"`
	MeasuredValue string `json:"measuredValue"`
}

type ForecastSqlResponse struct {
	Timestamp  string
	Pollutant  string
	Values     int
	Confidence float32
}
