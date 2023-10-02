package main

import (
	"iudx_domain_specific_apis/routes"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	routes.AirQuality(router)
	router.Run("localhost:9000")
}
