package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	cors "github.com/rs/cors/wrapper/gin"

	"iudx_domain_specific_apis/pkg/configs"
	"iudx_domain_specific_apis/pkg/db"
	"iudx_domain_specific_apis/pkg/logger"
	"iudx_domain_specific_apis/pkg/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Info.Println("Failed to load env file")
	}

	configs.Initialize()

	db.Init()

	router := gin.Default()

	// ? CORS Enabled
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	routes.AirQuality(router)

	port := configs.GetAPIPort()

	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error.Fatalf("listen: %s\n", err)
		}
	}()

	logger.Info.Printf("Starting server at %s", port)

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info.Println("Wait... Server shutting down......")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.Close(); err != nil {
		logger.Error.Println(err.Error())
	}
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()

	logger.Info.Println("Server exiting")
}
