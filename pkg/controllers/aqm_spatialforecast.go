package controllers

import (
	"iudx_domain_specific_apis/pkg/logger"
	"iudx_domain_specific_apis/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AQMSpatialForecastRequestBody struct {
	ForecastStart string `json:"forecastStart"`
	ForecastEnd   string `json:"forecastEnd"`
	GeoJson_Id    int32  `json:"geojson_id"`
}

type AQMSpatialForecasrController struct{}

var aqmSpatialForecastModel = new(models.AQMSpatialForecastModel)

func (ctrl AQMSpatialForecasrController) GetAQMSpatialForecast(c *gin.Context) {
	var reqBody AQMSpatialForecastRequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		logger.Error.Println(err.Error())
		return
	}

	aqmSpatialForecasts, err := aqmSpatialForecastModel.GetAQMSpatialForecasts(reqBody.ForecastStart, reqBody.ForecastEnd, reqBody.GeoJson_Id)

	if err != nil {
		logger.Error.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get forecasts"})
		return
	}

	c.JSON(http.StatusOK, aqmSpatialForecasts)
}
