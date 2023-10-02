package configs

import (
	"database/sql"
	"fmt"
	"iudx_domain_specific_apis/models"
	"iudx_domain_specific_apis/responses"
)

func ConnectDB() *sql.DB {

	envObj := GetEnv()

	connStr := "postgres://" + envObj.POSTGRES_USER + ":" + envObj.POSTGRES_PASSWORD + "@" + envObj.POSTGRES_HOST + "/" + envObj.POSTGRES_DB_NAME + "?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err.Error())
	}

	return db
}

var DB *sql.DB = ConnectDB()

func getForecast(client *sql.DB, queryString string, deviceId string) responses.ForecastResponse {

	rows, err := client.Query(queryString)

	if err != nil {
		fmt.Println(err.Error())
	}

	var reqBody models.ForecastRequestBody

	var timestamps []string
	var values []int
	var confidence []float32
	timeseries := responses.TimeseriesForecastResponse{}
	properties := responses.PropertiesForecastResponse{}
	properties.DeviceId = deviceId
	properties.MSE = 123
	properties.MeasuredValue = reqBody.MeasuredValue

	for rows.Next() {
		var forecast models.ForecastSqlResponse
		err := rows.Scan(&forecast.Timestamp, &forecast.Pollutant, &forecast.Values, &forecast.Confidence)

		if err != nil {
			fmt.Println(err.Error())
		}

		timestamps = append(timestamps, forecast.Timestamp)
		values = append(values, forecast.Values)
		confidence = append(confidence, float32(forecast.Confidence))
	}

	timeseries.Timestamps = timestamps
	timeseries.Values = values
	timeseries.Confidence = confidence

	respBody := responses.ForecastResponse{}
	respBody.Timeseries = timeseries
	respBody.Properties = properties

	rows.Close()
	client.Close()

	return respBody
}
