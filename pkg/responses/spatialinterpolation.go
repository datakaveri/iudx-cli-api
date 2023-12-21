package responses

import (
	"encoding/json"
	"iudx_domain_specific_apis/pkg/models"
	"sort"
	"time"
)

func FormatSpatialInterpolationResponse(spatialInterpolations []models.SpatialForecast, spatialInterpolationMinMax models.SpatialForecastMinMax, measuredValue string) SpatialForecastResponse {
	valueMap := make(map[string][]float32)
	geoJsonMap := make(map[string][]json.RawMessage)

	for _, spatialInterpolation := range spatialInterpolations {
		// ts := aqmSpatialForecast.ObservationDateTime.String()
		ts := spatialInterpolation.ObservationDateTime.Format("2006-01-02T15:04:05")
		valueMap[ts] = append(valueMap[ts], spatialInterpolation.PollutantVal)
		geoJsonMap[ts] = append(geoJsonMap[ts], json.RawMessage(spatialInterpolation.GeoJson))
	}

	var timestamps []string
	var values [][]float32
	var geoJson [][]json.RawMessage

	keys := make([]time.Time, 0, len(valueMap))

	for k := range valueMap {
		layout := "2006-01-02T15:04:05"
		time, _ := time.Parse(layout, k)
		keys = append(keys, time)
	}

	sort.Sort(ByDate(keys))

	for _, timestamp := range keys {
		newStamp := timestamp.Format("2006-01-02T15:04:05")
		timestamps = append(timestamps, newStamp)
		values = append(values, valueMap[newStamp])
		geoJson = append(geoJson, geoJsonMap[newStamp])
	}

	resultSpatialInterpolationResponse := SpatialForecastResponse{
		Timeseries: TimeseriesSpatialForecastResponse{
			Timestamps: timestamps,
			GeoJson:    geoJson,
			Values:     values,
		},
		Properties: PropertiesSpatialForecastResponse{
			MeasuredValue: measuredValue,
			MaxValue:      spatialInterpolationMinMax.Max,
			MinValue:      spatialInterpolationMinMax.Min,
			Average:       spatialInterpolationMinMax.Average,
			Stddev:        spatialInterpolationMinMax.Stddev,
		},
	}

	return resultSpatialInterpolationResponse
}
