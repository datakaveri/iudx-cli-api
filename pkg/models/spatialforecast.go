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
		WITH RankedData AS (
			select 
			kdmc_aqm_geojsons.geojson as geojson,
			kdmc_aqm_interpolation_forecast_data.observationdatetime as observationdatetime, 
			kdmc_aqm_interpolation_forecast_data.`+measuredValue+` as pollutant_val, 
			ROW_NUMBER() OVER (PARTITION BY kdmc_aqm_interpolation_forecast_data.hex_id, 
			kdmc_aqm_interpolation_forecast_data.observationdatetime 
			ORDER BY kdmc_aqm_interpolation_forecast_data.prediction_time DESC) AS RowNum
			from kdmc_aqm_interpolation_forecast_data
			inner join kdmc_aqm_geojsons on kdmc_aqm_geojsons.hex_id = kdmc_aqm_interpolation_forecast_data.hex_id
			where kdmc_aqm_interpolation_forecast_data.observationdatetime >= $1 
			and kdmc_aqm_interpolation_forecast_data.observationdatetime <= $2
		)
		SELECT geojson, observationdatetime, pollutant_val
		FROM RankedData
		WHERE RowNum = 1
		order by observationdatetime
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
