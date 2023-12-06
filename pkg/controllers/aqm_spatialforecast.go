package controllers

import (
	"iudx_domain_specific_apis/pkg/logger"
	"iudx_domain_specific_apis/pkg/models"
	"iudx_domain_specific_apis/pkg/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AQMSpatialForecastRequestBody struct {
	ForecastStart time.Time `json:"forecastStart" time_format:"2006-01-02T15:04:05Z07:00"`
	ForecastEnd   time.Time `json:"forecastEnd" time_format:"2006-01-02T15:04:05Z07:00"`
	MeasuredValue string    `json:"measuredValue"`
}

type AQMSpatialForecastController struct{}

var aqmSpatialForecastModel = new(models.AQMSpatialForecastModel)

func (ctrl AQMSpatialForecastController) GetAQMSpatialForecast(c *gin.Context) {
	var reqBody AQMSpatialForecastRequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		logger.Error.Println(err.Error())
		return
	}

	aqmSpatialForecasts, aqmSpatialForecastMinMax, err := aqmSpatialForecastModel.GetAQMSpatialForecasts(reqBody.ForecastStart, reqBody.ForecastEnd, reqBody.MeasuredValue)

	if err != nil {
		logger.Error.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get forecasts"})
		return
	}

	c.JSON(http.StatusOK, responses.FormatAQMSpatialForecastResponse(aqmSpatialForecasts, aqmSpatialForecastMinMax[0], reqBody.MeasuredValue))
}
