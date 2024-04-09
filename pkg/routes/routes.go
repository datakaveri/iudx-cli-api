package routes

import (
	"iudx_domain_specific_apis/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func AirQuality(router *gin.Engine) {

	// forecast := new(controllers.ForecastController)
	// router.POST("/airquality/forecast/:deviceID", middleware.APIKeyAuthMiddleware(), forecast.GetForecasts)

	// // spatialForecast := new(controllers.SpatialForecastController_LEGACY)
	// // router.POST("/airquality/spatialForecast", middleware.APIKeyAuthMiddleware(), spatialForecast.GetSpatialForecast_LEGACY)

	// spatialForecast := new(controllers.SpatialForecastController)
	// router.POST("/airquality/spatialForecast", middleware.APIKeyAuthMiddleware(), spatialForecast.GetSpatialForecast)

	// spatialInterpolation := new(controllers.SpatialInterpolationController)
	// router.POST("/airquality/spatialInterpolation", middleware.APIKeyAuthMiddleware(), spatialInterpolation.GetSpatialInterpolation)

	spatialForecast := new(controllers.SpatialForecastController)
	router.POST("/airquality/pimpriSpatialForecast", spatialForecast.GetSpatialForecast)

	spatialInterpolation := new(controllers.SpatialInterpolationController)
	router.POST("/airquality/pimpriSpatialInterpolation", spatialInterpolation.GetSpatialInterpolation)

	landing := new(controllers.LandingController)
	router.GET("/", landing.Landing)
}
