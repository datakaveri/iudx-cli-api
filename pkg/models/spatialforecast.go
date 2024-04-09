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
			pimpri_aqm_geojsons.geojson as geojson,
			pimpri_aqm_interpolation_forecast_data.observationdatetime as observationdatetime, 
			pimpri_aqm_interpolation_forecast_data.`+measuredValue+` as pollutant_val, 
			ROW_NUMBER() OVER (PARTITION BY pimpri_aqm_interpolation_forecast_data.hex_id, 
			pimpri_aqm_interpolation_forecast_data.observationdatetime 
			ORDER BY pimpri_aqm_interpolation_forecast_data.prediction_time DESC) AS RowNum
			from pimpri_aqm_interpolation_forecast_data
			inner join pimpri_aqm_geojsons on pimpri_aqm_geojsons.hex_id = pimpri_aqm_interpolation_forecast_data.hex_id
			where pimpri_aqm_interpolation_forecast_data.observationdatetime >= $1 
			and pimpri_aqm_interpolation_forecast_data.observationdatetime <= $2
		)
		SELECT geojson, observationdatetime, pollutant_val
		FROM RankedData
		WHERE RowNum = 1
		order by observationdatetime
	`, startTime, endTime)

	_, err = db.GetDB().Select(&spatialForecastMinMax, `
		select min(pimpri_aqm_interpolation_forecast_data.`+measuredValue+`), 
		max(pimpri_aqm_interpolation_forecast_data.`+measuredValue+`) ,
		average(stats_agg(pimpri_aqm_interpolation_forecast_data.`+measuredValue+`)),
		stddev(stats_agg(pimpri_aqm_interpolation_forecast_data.`+measuredValue+`), 'pop')
		from pimpri_aqm_interpolation_forecast_data
	`)

	return spatialForecast, spatialForecastMinMax, err
}
