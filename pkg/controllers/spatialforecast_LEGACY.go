package controllers

import (
	"iudx_domain_specific_apis/pkg/logger"
	"iudx_domain_specific_apis/pkg/models"
	"iudx_domain_specific_apis/pkg/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SpatialForecastRequestBody_LEGACY struct {
	ForecastStart time.Time `json:"forecastStart" time_format:"2006-01-02T15:04:05Z07:00"`
	ForecastEnd   time.Time `json:"forecastEnd" time_format:"2006-01-02T15:04:05Z07:00"`
	MeasuredValue string    `json:"measuredValue"`
}

type SpatialForecastController_LEGACY struct{}

var spatialForecastModel_LEGACY = new(models.SpatialForecastModel_LEGACY)

func (ctrl SpatialForecastController_LEGACY) GetSpatialForecast_LEGACY(c *gin.Context) {

	var reqBody SpatialForecastRequestBody_LEGACY
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		logger.Error.Println(err.Error())
		return
	}

	spatialForecasts, err := spatialForecastModel_LEGACY.GetSpatialForecasts_LEGACY(reqBody.ForecastStart, reqBody.ForecastEnd)
	if err != nil {
		logger.Error.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get forecasts"})
		return
	}

	c.JSON(http.StatusOK, responses.FormatSpatialForecastResponse_LEGACY(spatialForecasts))
}
