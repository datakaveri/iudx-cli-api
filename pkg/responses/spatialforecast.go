package responses

import (
	"encoding/json"
	"iudx_domain_specific_apis/pkg/models"
)

type TimeseriesSpatialForecastResponse struct {
	Timestamps []string            `json:"timestamps"`
	GeoJson    [][]json.RawMessage `json:"geojson"`
	Values     [][]float32         `json:"values"`
}

type PropertiesSpatialForecastResponse struct {
	MeasuredValue string `json:"measuredValue"`
}

type SpatialForecastResponse struct {
	Timeseries TimeseriesSpatialForecastResponse `json:"timeseries"`
	Properties PropertiesSpatialForecastResponse `json:"properties"`
}

func FormatSpatialForecastResponse(spatialForecasts []models.SpatialForecast) SpatialForecastResponse {
	valueMap := make(map[string][]float32)
	geoJsonMap := make(map[string][]json.RawMessage)

	for _, spatialForecast := range spatialForecasts {
		ts := spatialForecast.ObservationDateTime.Format("2006-01-02T15:04:05+00:00")
		valueMap[ts] = append(valueMap[ts], spatialForecast.PollutantVal)

		geoJsonMap[ts] = append(geoJsonMap[ts], spatialForecast.GeoJson)
	}

	var timestamps []string
	var values [][]float32
	var geoJson [][]json.RawMessage

	// Populate the slices from the maps
	for timestamp, val := range valueMap {
		timestamps = append(timestamps, timestamp)
		values = append(values, val)
	}

	for _, geo := range geoJsonMap {
		geoJson = append(geoJson, geo)
	}

	resultSpatialForecastResponse := SpatialForecastResponse{
		Timeseries: TimeseriesSpatialForecastResponse{
			Timestamps: timestamps,
			Values:     values,
			GeoJson:    geoJson,
		},
		Properties: PropertiesSpatialForecastResponse{
			MeasuredValue: "co2",
		},
	}

	return resultSpatialForecastResponse
}
