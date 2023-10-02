package responses

type TimeseriesForecastResponse struct {
	Timestamps []string  `json:"timestamps"`
	Values     []int     `json:"values"`
	Confidence []float32 `json:"confidence"`
}

type PropertiesForecastResponse struct {
	DeviceId      string `json:"deviceId"`
	MeasuredValue string `json:"measuredValue"`
	MSE           int    `json:"MSE"`
}

type ForecastResponse struct {
	Timeseries TimeseriesForecastResponse `json:"timeseries"`
	Properties PropertiesForecastResponse `json:"properties"`
}
