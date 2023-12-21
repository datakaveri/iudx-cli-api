package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LandingController struct{}

func (ctrl LandingController) Landing(c *gin.Context) {
	c.HTML(http.StatusOK, "landing.html", gin.H{
		"title": "Landing Page",
	})
}
