package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SpatialForecast() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("spatialForecast")
	}
}
