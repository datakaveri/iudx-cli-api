package controllers

import (
	"iudx_domain_specific_apis/pkg/logger"
	"iudx_domain_specific_apis/pkg/models"
	"iudx_domain_specific_apis/pkg/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SpatialInterpolationRequestBody struct {
	Start         time.Time `json:"start" time_format:"2006-01-02T15:04:05Z07:00"`
	End           time.Time `json:"end" time_format:"2006-01-02T15:04:05Z07:00"`
	MeasuredValue string    `json:"measuredValue"`
}

type SpatialInterpolationController struct{}

var spatialInterpolationModel = new(models.SpatialInterpolationModel)

func (ctrl SpatialInterpolationController) GetSpatialInterpolation(c *gin.Context) {
	var reqBody SpatialInterpolationRequestBody

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		logger.Error.Println(err.Error())
		return
	}

	spatialInterpolations, spatialInterpolationsMinMax, err := spatialInterpolationModel.GetSpatialInterpolations(reqBody.Start, reqBody.End, reqBody.MeasuredValue)

	if err != nil {
		logger.Error.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Could not get forecasts"})
		return
	}

	c.JSON(http.StatusOK, responses.FormatSpatialInterpolationResponse(spatialInterpolations, spatialInterpolationsMinMax[0], reqBody.MeasuredValue))
}
