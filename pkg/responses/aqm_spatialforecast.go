package responses

import (
	"encoding/json"
	"fmt"
	"iudx_domain_specific_apis/pkg/models"
)

type TimeseriesAQMSpatialForecastResponse struct {
	Timestamps []string            `json:"timestamps"`
	GeoJson    [][]json.RawMessage `json:"geojson"`
	Values     [][]float32         `json:"values"`
}

type PropertiesAQMSpatialForecastResponse struct {
	MeasuredValue string `json:"measuredValue"`
}

type AQMSpatialForecastResponse struct {
	Timeseries TimeseriesAQMSpatialForecastResponse `json:"timeseries"`
	Properties PropertiesAQMSpatialForecastResponse `json:"properties"`
}

func FormatAQMSpatialForecastResponse(aqmSpatialForecasts []models.AQMSpatialForecast) AQMSpatialForecastResponse {
	valueMap := make(map[string][]float32)
	geoJsonMap := make(map[string][]json.RawMessage)

	for _, aqmSpatialForecast := range aqmSpatialForecasts {
		// ts := aqmSpatialForecast.ObservationDateTime.String()
		ts := aqmSpatialForecast.ObservationDateTime.Format("2006-01-02T15:04:05+00:00")
		valueMap[ts] = append(valueMap[ts], aqmSpatialForecast.PollutantVal)
		geoJsonMap[ts] = append(geoJsonMap[ts], json.RawMessage(aqmSpatialForecast.GeoJson))
	}

	var timestamps []string
	var values [][]float32
	var geoJson [][]json.RawMessage

	for timestamp, val := range valueMap {
		timestamps = append(timestamps, timestamp)
		values = append(values, val)
	}

	for _, geo := range geoJsonMap {
		geoJson = append(geoJson, geo)
	}

	fmt.Println(timestamps)

	resultAQMSpatialForecastResponse := AQMSpatialForecastResponse{
		Timeseries: TimeseriesAQMSpatialForecastResponse{
			Timestamps: timestamps,
			GeoJson:    geoJson,
			Values:     values,
		},
		Properties: PropertiesAQMSpatialForecastResponse{
			MeasuredValue: "co2",
		},
	}

	return resultAQMSpatialForecastResponse
}
