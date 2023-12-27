package responses

import (
	"encoding/json"
	"iudx_domain_specific_apis/pkg/models"
	"sort"
	"time"
)

type TimeseriesSpatialForecastResponse struct {
	Timestamps []string            `json:"timestamps"`
	GeoJson    [][]json.RawMessage `json:"geojson"`
	Values     [][]float32         `json:"values"`
}

type PropertiesSpatialForecastResponse struct {
	MeasuredValue string  `json:"measuredValue"`
	MinValue      float32 `json:"min"`
	MaxValue      float32 `json:"max"`
	Average       float32 `json:"average"`
	Stddev        float32 `json:"stddev"`
}

type SpatialForecastResponse struct {
	Timeseries TimeseriesSpatialForecastResponse `json:"timeseries"`
	Properties PropertiesSpatialForecastResponse `json:"properties"`
}

type ByDate []time.Time

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Before(a[j]) }

func FormatSpatialForecastResponse(spatialForecasts []models.SpatialForecast, spatialForecastMinMax models.SpatialForecastMinMax, measuredValue string) SpatialForecastResponse {
	valueMap := make(map[string][]float32)
	geoJsonMap := make(map[string][]json.RawMessage)

	for _, spatialForecast := range spatialForecasts {
		// ts := aqmSpatialForecast.ObservationDateTime.String()
		ts := spatialForecast.ObservationDateTime.Format("2006-01-02T15:04:05")
		valueMap[ts] = append(valueMap[ts], spatialForecast.PollutantVal)
		geoJsonMap[ts] = append(geoJsonMap[ts], json.RawMessage(spatialForecast.GeoJson))
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

	resultSpatialForecastResponse := SpatialForecastResponse{
		Timeseries: TimeseriesSpatialForecastResponse{
			Timestamps: timestamps,
			GeoJson:    geoJson,
			Values:     values,
		},
		Properties: PropertiesSpatialForecastResponse{
			MeasuredValue: measuredValue,
			MaxValue:      spatialForecastMinMax.Max,
			MinValue:      spatialForecastMinMax.Min,
			Average:       spatialForecastMinMax.Average,
			Stddev:        spatialForecastMinMax.Stddev,
		},
	}

	return resultSpatialForecastResponse
}
