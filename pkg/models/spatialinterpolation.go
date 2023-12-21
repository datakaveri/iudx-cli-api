package models

import (
	"iudx_domain_specific_apis/pkg/db"
	"time"
)

type SpatialInterpolationModel struct{}

func (m SpatialInterpolationModel) GetSpatialInterpolations(startTime time.Time, endTime time.Time, measuredValue string) (spatialInterpolation []SpatialForecast, spatialInterpolationMinMax []SpatialForecastMinMax, err error) {

	_, err = db.GetDB().Select(&spatialInterpolation, `
		select kdmc_aqm_geojsons.geojson,
		kdmc_aqm_interpolation_actual_data.observationdatetime, kdmc_aqm_interpolation_actual_data.`+measuredValue+` as pollutant_val
		from kdmc_aqm_geojsons
		inner join kdmc_aqm_interpolation_actual_data on kdmc_aqm_geojsons.hex_id = kdmc_aqm_interpolation_actual_data.hex_id
		where observationdatetime >= $1 and observationdatetime <= $2
		order by kdmc_aqm_interpolation_actual_data.observationdatetime
	`, startTime, endTime)

	_, err = db.GetDB().Select(&spatialInterpolationMinMax, `
		select min(kdmc_aqm_interpolation_actual_data.`+measuredValue+`), 
		max(kdmc_aqm_interpolation_actual_data.`+measuredValue+`) ,
		average(stats_agg(kdmc_aqm_interpolation_actual_data.`+measuredValue+`)),
		stddev(stats_agg(kdmc_aqm_interpolation_actual_data.`+measuredValue+`), 'pop')
		from kdmc_aqm_interpolation_actual_data
	`)

	return spatialInterpolation, spatialInterpolationMinMax, err
}
