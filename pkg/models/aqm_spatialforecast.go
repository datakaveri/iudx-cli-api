package models

import (
	"iudx_domain_specific_apis/pkg/db"
)

type AQMSpatialForecast struct {
	GeoJson_Id          int32   `db:"geojson_id" json:"geojson_id"`
	GeoJson             string  `db:"geojson" json:"geojson"`
	ObservationDateTime string  `db:"observationdatetime" json:"observationdatetime"`
	PollutantVal        float32 `db:"pollutant_val" json:"pollutant_val"`
}

type AQMSpatialForecastModel struct{}

func (m AQMSpatialForecastModel) GetAQMSpatialForecasts(startTime string, endTime string, geojson_id int32) (aqmSpatialForecast []AQMSpatialForecast, err error) {

	_, err = db.GetDB().Select(&aqmSpatialForecast, `
		select aqm_geojson.geojson_id, aqm_geojson.geojson,
		aqm_forecast.observationdatetime, aqm_forecast.pollutant_val
		from aqm_geojson
		inner join aqm_forecast on aqm_geojson.geojson_id = aqm_forecast.geojson_id
		where observationdatetime >= $1 and observationdatetime <= $2 and aqm_geojson.geojson_id = $3
	`, startTime, endTime, geojson_id)
	return aqmSpatialForecast, err
}
