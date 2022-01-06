package main

import (
	"github.com/msfidelis/health-api/controllers/calculator"
	"github.com/msfidelis/health-api/controllers/healthcheck"
	"github.com/msfidelis/health-api/controllers/liveness"
	"github.com/msfidelis/health-api/controllers/readiness"
	"github.com/msfidelis/health-api/controllers/version"
	"github.com/msfidelis/health-api/pkg/memory_cache"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	chaos "github.com/msfidelis/gin-chaos-monkey"

	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"

	_ "github.com/msfidelis/health-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	router := gin.New()

	// Logger
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	//Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Memory Cache Singleton
	c := memory_cache.GetInstance()

	// Readiness Probe Mock Config
	probe_time_raw := os.Getenv("READINESS_PROBE_MOCK_TIME_IN_SECONDS")
	if probe_time_raw == "" {
		probe_time_raw = "5"
	}
	probe_time, err := strconv.ParseUint(probe_time_raw, 10, 64)
	if err != nil {
		fmt.Println("Environment variable READINESS_PROBE_MOCK_TIME_IN_SECONDS conversion error", err)
	}
	c.Set("readiness.ok", "false", time.Duration(probe_time)*time.Second)

	// Prometheus Exporter Config
	p := ginprom.New(
		ginprom.Engine(router),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)

	//Middlewares
	router.Use(p.Instrument())
	router.Use(gin.Recovery())
	router.Use(chaos.Load())

	// Healthcheck Router
	router.GET("/healthcheck", healthcheck.Ok)

	// Version Router
	router.GET("/version", version.Get)

	// Liveness and Readiness
	router.GET("/liveness", liveness.Ok)
	router.GET("/readiness", readiness.Ok)

	// Health Calculator
	router.POST("/calculator", calculator.Post)

	router.Run()
}
