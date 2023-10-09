package routes

import (
	"iudx_domain_specific_apis/pkg/controllers"
	"iudx_domain_specific_apis/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func AirQuality(router *gin.Engine) {

	forecast := new(controllers.ForecastController)
	router.POST("/airquality/forecast/:deviceID", middleware.APIKeyAuthMiddleware(), forecast.GetForecasts)

	spatialForecast := new(controllers.SpatialForecastController)
	router.POST("/airquality/spatialForecast", middleware.APIKeyAuthMiddleware(), spatialForecast.GetSpatialForecast)
}
