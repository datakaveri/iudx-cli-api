package routes

import (
	"iudx_domain_specific_apis/controllers"

	"github.com/gin-gonic/gin"
)

func AirQuality(router *gin.Engine) {
	router.POST("/airquality/forecast", controllers.Forecast())
	router.POST("/airquality/spatialForecast", controllers.SpatialForecast())
}
