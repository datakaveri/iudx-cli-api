package models

import (
	"encoding/json"
	"iudx_domain_specific_apis/pkg/db"
	"time"
)

type SpatialForecast struct {
	GeoJson             json.RawMessage `db:"geojson" json:"geojson"`
	PollutantVal        float32         `db:"pollutant_val" json:"pollutant_val"`
	ObservationDateTime time.Time       `db:"observationdatetime" json:"observationdatetime"`
}

type SpatialForecastModel struct{}

func (m SpatialForecastModel) GetSpatialForecasts(startTime time.Time, endTime time.Time) (spatialforecast []SpatialForecast, err error) {
	_, err = db.GetDB().Select(&spatialforecast, "SELECT * FROM spatial_forecast WHERE observationdatetime BETWEEN $1 AND $2 ", startTime, endTime)
	return spatialforecast, err
}
