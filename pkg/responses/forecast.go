package responses

import (
	"iudx_domain_specific_apis/pkg/models"
	"time"
)

type TimeseriesForecastResponse struct {
	Timestamps []time.Time `json:"timestamps"`
	Values     []float32   `json:"values"`
}

type PropertiesForecastResponse struct {
	DeviceId      string `json:"deviceId"`
	MeasuredValue string `json:"measuredValue"`
}

type ForecastResponse struct {
	Timeseries TimeseriesForecastResponse `json:"timeseries"`
	Properties PropertiesForecastResponse `json:"properties"`
}

func FormatForecastResponse(forecasts []models.Forecast) ForecastResponse {
	var timeseriesTimes []time.Time
	var value []float32

	for _, forecast := range forecasts {
		timeseriesTimes = append(timeseriesTimes, forecast.ObservationDateTime)
		value = append(value, forecast.Co2)

	}

	resultForecastResponse := ForecastResponse{
		Timeseries: TimeseriesForecastResponse{
			Timestamps: timeseriesTimes,
			Values:     value,
		},
		Properties: PropertiesForecastResponse{
			DeviceId:      forecasts[0].DeviceId,
			MeasuredValue: "co2",
		},
	}

	return resultForecastResponse
}
