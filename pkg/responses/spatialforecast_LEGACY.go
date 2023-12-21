package responses

import (
	"encoding/json"
	"iudx_domain_specific_apis/pkg/models"
)

type TimeseriesSpatialForecastResponse_LEGACY struct {
	Timestamps []string            `json:"timestamps"`
	GeoJson    [][]json.RawMessage `json:"geojson"`
	Values     [][]float32         `json:"values"`
}

type PropertiesSpatialForecastResponse_LEGACY struct {
	MeasuredValue string `json:"measuredValue"`
}

type SpatialForecastResponse_LEGACY struct {
	Timeseries TimeseriesSpatialForecastResponse_LEGACY `json:"timeseries"`
	Properties PropertiesSpatialForecastResponse_LEGACY `json:"properties"`
}

func FormatSpatialForecastResponse_LEGACY(spatialForecasts_LEGACY []models.SpatialForecast_LEGACY) SpatialForecastResponse_LEGACY {
	valueMap := make(map[string][]float32)
	geoJsonMap := make(map[string][]json.RawMessage)

	for _, spatialForecast_LEGACY := range spatialForecasts_LEGACY {
		ts := spatialForecast_LEGACY.ObservationDateTime.Format("2006-01-02T15:04:05+00:00")
		valueMap[ts] = append(valueMap[ts], spatialForecast_LEGACY.PollutantVal)

		geoJsonMap[ts] = append(geoJsonMap[ts], spatialForecast_LEGACY.GeoJson)
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

	resultSpatialForecastResponse_LEGACY := SpatialForecastResponse_LEGACY{
		Timeseries: TimeseriesSpatialForecastResponse_LEGACY{
			Timestamps: timestamps,
			Values:     values,
			GeoJson:    geoJson,
		},
		Properties: PropertiesSpatialForecastResponse_LEGACY{
			MeasuredValue: "co2",
		},
	}

	return resultSpatialForecastResponse_LEGACY
}
