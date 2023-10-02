package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"iudx_domain_specific_apis/models"
	"iudx_domain_specific_apis/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Forecast_LEGACY() gin.HandlerFunc {
	return func(c *gin.Context) {
		deviceId := c.Param("id")
		fmt.Println("forecast")
		// var cache = memcache.New("localhost:11211")

		// val, cacheErr := cache.Get(deviceId)

		// if cacheErr != nil {
		// 	fmt.Println(cacheErr.Error())
		// }

		// if val != nil {
		// 	fmt.Println(string(val.Value))
		// 	respBody := responses.ForecastResponse{}
		// 	json.Unmarshal([]byte(val.Value), &respBody)
		// 	c.IndentedJSON(http.StatusOK, respBody)
		// } else {

		connStr := "postgres://postgres:password@localhost/z?sslmode=disable"
		db, err := sql.Open("postgres", connStr)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		}

		var reqBody models.ForecastRequestBody
		if err := c.BindJSON(&reqBody); err != nil {
			fmt.Println(err.Error())
			return
		}

		var timestamps []string
		var values []int
		var confidence []float32
		timeseries := responses.TimeseriesForecastResponse{}
		properties := responses.PropertiesForecastResponse{}
		properties.DeviceId = deviceId
		properties.MSE = 123
		properties.MeasuredValue = reqBody.MeasuredValue

		rows, err := db.Query(`select * from Forecast srt where pollutant=$1 and timestamp>$2 and timestamp<$3`, reqBody.MeasuredValue, reqBody.ForecastStart, reqBody.ForecastEnd)

		for rows.Next() {
			var forecast models.ForecastSqlResponse
			err := rows.Scan(&forecast.Timestamp, &forecast.Pollutant, &forecast.Values, &forecast.Confidence)

			if err != nil {
				c.IndentedJSON(http.StatusBadRequest, err.Error())
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
		db.Close()

		jsonResp, jsonErr := json.Marshal(respBody)

		if jsonErr != nil {
			fmt.Println(jsonErr.Error())
		}

		verifyCache(deviceId, string(jsonResp))

		c.IndentedJSON(http.StatusOK, respBody)
		// }
	}
}
