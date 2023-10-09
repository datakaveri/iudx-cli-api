package models

import (
	"iudx_domain_specific_apis/pkg/db"
	"time"
)

type Forecast struct {
	DeviceId            string    `db:"device_id" json:"device_id"`
	Co2                 float32   `db:"co2" json:"co2"`
	ObservationDateTime time.Time `db:"observationdatetime" json:"observationdatetime"`
}

type ForecastModel struct{}

func (m ForecastModel) GetForecasts(deviceID string, startTime time.Time, endTime time.Time) (forecast []Forecast, err error) {
	_, err = db.GetDB().Select(&forecast, "SELECT device_id, observationdatetime, co2 FROM forecast WHERE device_id = $1 AND observationdatetime BETWEEN $2 AND $3 ", deviceID, startTime, endTime)
	return forecast, err
}
