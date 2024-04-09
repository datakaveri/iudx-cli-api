package models

import (
	"iudx_domain_specific_apis/pkg/db"
	"time"
)

type SpatialInterpolationModel struct{}

func (m SpatialInterpolationModel) GetSpatialInterpolations(startTime time.Time, endTime time.Time, measuredValue string) (spatialInterpolation []SpatialForecast, spatialInterpolationMinMax []SpatialForecastMinMax, err error) {

	_, err = db.GetDB().Select(&spatialInterpolation, `
		select pimpri_aqm_geojsons.geojson,
		pimpri_aqm_interpolation_actual_data.observationdatetime, pimpri_aqm_interpolation_actual_data.`+measuredValue+` as pollutant_val
		from pimpri_aqm_geojsons
		inner join pimpri_aqm_interpolation_actual_data on pimpri_aqm_geojsons.hex_id = pimpri_aqm_interpolation_actual_data.hex_id
		where observationdatetime >= $1 and observationdatetime <= $2
		order by pimpri_aqm_interpolation_actual_data.observationdatetime
	`, startTime, endTime)

	_, err = db.GetDB().Select(&spatialInterpolationMinMax, `
		select min(pimpri_aqm_interpolation_actual_data.`+measuredValue+`), 
		max(pimpri_aqm_interpolation_actual_data.`+measuredValue+`) ,
		average(stats_agg(pimpri_aqm_interpolation_actual_data.`+measuredValue+`)),
		stddev(stats_agg(pimpri_aqm_interpolation_actual_data.`+measuredValue+`), 'pop')
		from pimpri_aqm_interpolation_actual_data
	`)

	return spatialInterpolation, spatialInterpolationMinMax, err
}
