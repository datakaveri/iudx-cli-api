package models

import (
	"encoding/json"
	"iudx_domain_specific_apis/pkg/db"
	"time"
)

type SpatialForecast_LEGACY struct {
	GeoJson             json.RawMessage `db:"geojson" json:"geojson"`
	PollutantVal        float32         `db:"pollutant_val" json:"pollutant_val"`
	ObservationDateTime time.Time       `db:"observationdatetime" json:"observationdatetime"`
}

type SpatialForecastModel_LEGACY struct{}

func (m SpatialForecastModel_LEGACY) GetSpatialForecasts_LEGACY(startTime time.Time, endTime time.Time) (spatialforecast_LEGACY []SpatialForecast_LEGACY, err error) {
	_, err = db.GetDB().Select(&spatialforecast_LEGACY, "SELECT * FROM spatial_forecast WHERE observationdatetime BETWEEN $1 AND $2 ", startTime, endTime)
	return spatialforecast_LEGACY, err
}
