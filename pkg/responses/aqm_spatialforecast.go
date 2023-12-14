package responses

import (
	"encoding/json"
	"iudx_domain_specific_apis/pkg/models"
	"sort"
	"time"
)

type TimeseriesAQMSpatialForecastResponse struct {
	Timestamps []string            `json:"timestamps"`
	GeoJson    [][]json.RawMessage `json:"geojson"`
	Values     [][]float32         `json:"values"`
}

type PropertiesAQMSpatialForecastResponse struct {
	MeasuredValue string  `json:"measuredValue"`
	MinValue      float32 `json:"min"`
	MaxValue      float32 `json:"max"`
	Average       float32 `json:"average"`
	Stddev        float32 `json:"stddev"`
}

type AQMSpatialForecastResponse struct {
	Timeseries TimeseriesAQMSpatialForecastResponse `json:"timeseries"`
	Properties PropertiesAQMSpatialForecastResponse `json:"properties"`
}

type ByDate []time.Time

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Before(a[j]) }

func FormatAQMSpatialForecastResponse(aqmSpatialForecasts []models.AQMSpatialForecast, aqmSpatialForecastMinMax models.AQMSpatialForecastMinMax, measuredValue string) AQMSpatialForecastResponse {
	valueMap := make(map[string][]float32)
	geoJsonMap := make(map[string][]json.RawMessage)

	for _, aqmSpatialForecast := range aqmSpatialForecasts {
		// ts := aqmSpatialForecast.ObservationDateTime.String()
		ts := aqmSpatialForecast.ObservationDateTime.Format("2006-01-02T15:04:05")
		valueMap[ts] = append(valueMap[ts], aqmSpatialForecast.PollutantVal)
		geoJsonMap[ts] = append(geoJsonMap[ts], json.RawMessage(aqmSpatialForecast.GeoJson))
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

	resultAQMSpatialForecastResponse := AQMSpatialForecastResponse{
		Timeseries: TimeseriesAQMSpatialForecastResponse{
			Timestamps: timestamps,
			GeoJson:    geoJson,
			Values:     values,
		},
		Properties: PropertiesAQMSpatialForecastResponse{
			MeasuredValue: measuredValue,
			MaxValue:      aqmSpatialForecastMinMax.Max,
			MinValue:      aqmSpatialForecastMinMax.Min,
			Average:       aqmSpatialForecastMinMax.Average,
			Stddev:        aqmSpatialForecastMinMax.Stddev,
		},
	}

	return resultAQMSpatialForecastResponse
}
