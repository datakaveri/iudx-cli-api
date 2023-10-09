package controllers

import (
	"iudx_domain_specific_apis/pkg/logger"
	"iudx_domain_specific_apis/pkg/models"
	"iudx_domain_specific_apis/pkg/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ForecastRequestBody struct {
	ForecastStart time.Time `json:"forecastStart" time_format:"2006-01-02T15:04:05Z07:00"`
	ForecastEnd   time.Time `json:"forecastEnd" time_format:"2006-01-02T15:04:05Z07:00"`
	MeasuredValue string    `json:"measuredValue"`
}

type ForecastController struct{}

var forecastModel = new(models.ForecastModel)

func (ctrl ForecastController) GetForecasts(c *gin.Context) {
	deviceID := c.Param("deviceID")

	// check if deviceId is empty
	if deviceID == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	var reqBody ForecastRequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		logger.Error.Println(err.Error())
		return
	}

	forecasts, err := forecastModel.GetForecasts(deviceID, reqBody.ForecastStart, reqBody.ForecastEnd)
	if err != nil {
		logger.Error.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get forecasts"})
		return
	}

	c.JSON(http.StatusOK, responses.FormatForecastResponse(forecasts))

}
