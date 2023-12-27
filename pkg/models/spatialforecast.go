package models

import (
	"iudx_domain_specific_apis/pkg/db"
	"time"
)

type SpatialForecast struct {
	GeoJson             string    `db:"geojson" json:"geojson"`
	ObservationDateTime time.Time `db:"observationdatetime" json:"observationdatetime"`
	PollutantVal        float32   `db:"pollutant_val" json:"pollutant_val"`
}

type SpatialForecastMinMax struct {
	Min     float32 `db:"min" json:"min"`
	Max     float32 `db:"max" json:"max"`
	Average float32 `db:"average" json:"average"`
	Stddev  float32 `db:"stddev" json:"stddev"`
}

type SpatialForecastModel struct{}

func (m SpatialForecastModel) GetSpatialForecasts(startTime time.Time, endTime time.Time, measuredValue string) (spatialForecast []SpatialForecast, spatialForecastMinMax []SpatialForecastMinMax, err error) {

	_, err = db.GetDB().Select(&spatialForecast, `
		select kdmc_aqm_geojsons.geojson,
		kdmc_aqm_interpolation_forecast_data.observationdatetime, kdmc_aqm_interpolation_forecast_data.`+measuredValue+` as pollutant_val
		from kdmc_aqm_geojsons
		inner join kdmc_aqm_interpolation_forecast_data on kdmc_aqm_geojsons.hex_id = kdmc_aqm_interpolation_forecast_data.hex_id
		where observationdatetime >= $1 and observationdatetime <= $2
		order by kdmc_aqm_interpolation_forecast_data.observationdatetime
	`, startTime, endTime)

	_, err = db.GetDB().Select(&spatialForecastMinMax, `
		select min(kdmc_aqm_interpolation_forecast_data.`+measuredValue+`), 
		max(kdmc_aqm_interpolation_forecast_data.`+measuredValue+`) ,
		average(stats_agg(kdmc_aqm_interpolation_forecast_data.`+measuredValue+`)),
		stddev(stats_agg(kdmc_aqm_interpolation_forecast_data.`+measuredValue+`), 'pop')
		from kdmc_aqm_interpolation_forecast_data
	`)

	return spatialForecast, spatialForecastMinMax, err
}
