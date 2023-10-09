package controllers

import (
	"iudx_domain_specific_apis/pkg/logger"
	"iudx_domain_specific_apis/pkg/models"
	"iudx_domain_specific_apis/pkg/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SpatialForecastRequestBody struct {
	ForecastStart time.Time `json:"forecastStart" time_format:"2006-01-02T15:04:05Z07:00"`
	ForecastEnd   time.Time `json:"forecastEnd" time_format:"2006-01-02T15:04:05Z07:00"`
	MeasuredValue string    `json:"measuredValue"`
}

type SpatialForecastController struct{}

var spatialForecastModel = new(models.SpatialForecastModel)

func (ctrl SpatialForecastController) GetSpatialForecast(c *gin.Context) {

	var reqBody SpatialForecastRequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		logger.Error.Println(err.Error())
		return
	}

	spatialForecasts, err := spatialForecastModel.GetSpatialForecasts(reqBody.ForecastStart, reqBody.ForecastEnd)
	if err != nil {
		logger.Error.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get forecasts"})
		return
	}

	c.JSON(http.StatusOK, responses.FormatSpatialForecastResponse(spatialForecasts))
}
