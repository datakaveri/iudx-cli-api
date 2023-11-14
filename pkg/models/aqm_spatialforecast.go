package models

import (
	"iudx_domain_specific_apis/pkg/db"
	"time"
)

type AQMSpatialForecast struct {
	GeoJson             string    `db:"geojson" json:"geojson"`
	ObservationDateTime time.Time `db:"observationdatetime" json:"observationdatetime"`
	PollutantVal        float32   `db:"co2" json:"pollutant_val"`
}

type AQMSpatialForecastModel struct{}

func (m AQMSpatialForecastModel) GetAQMSpatialForecasts(startTime time.Time, endTime time.Time, measuredValue string) (aqmSpatialForecast []AQMSpatialForecast, err error) {

	_, err = db.GetDB().Select(&aqmSpatialForecast, `
		select kdmc_aqm_geojsons.geojson,
		kdmc_aqm_interpolation_actual_data.observationdatetime, kdmc_aqm_interpolation_actual_data.co2
		from kdmc_aqm_geojsons
		inner join kdmc_aqm_interpolation_actual_data on kdmc_aqm_geojsons.hex_id = kdmc_aqm_interpolation_actual_data.hex_id
		where observationdatetime >= $1 and observationdatetime <= $2
		order by kdmc_aqm_interpolation_actual_data.observationdatetime
		`, startTime, endTime)
	return aqmSpatialForecast, err
}
