package main

import (
    "log"
    "os"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/jkloris/xkapustaj-webapi/api"
    "github.com/jkloris/xkapustaj-webapi/internal/xkapustaj_wl"
    "github.com/jkloris/xkapustaj-webapi/internal/db_service"
    "context"
    "time"
    "github.com/gin-contrib/cors"
)

func main() {
    log.Printf("Server started")
    port := os.Getenv("AMBULANCE_API_PORT ")
    if port == "" {
        port = "8080"
    }
    environment := os.Getenv("AMBULANCE_API_ENVIRONMENT")
    if !strings.EqualFold(environment, "production") { // case insensitive comparison
        gin.SetMode(gin.DebugMode)
    }
    engine := gin.New()
    engine.Use(gin.Recovery())
    corsMiddleware := cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH"},
        AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
        ExposeHeaders:    []string{""},
        AllowCredentials: false,
        MaxAge: 12 * time.Hour,
    })
    engine.Use(corsMiddleware)

    // setup context update  middleware
    dbService := db_service.NewMongoService[xkapustaj_wl.Hospital](db_service.MongoServiceConfig{})
    defer dbService.Disconnect(context.Background())
    engine.Use(func(ctx *gin.Context) {
        ctx.Set("db_service", dbService)
        ctx.Next()
    })

    // request routings
    xkapustaj_wl.AddRoutes(engine)
    engine.GET("/openapi", api.HandleOpenApi)
    engine.Run(":" + port)
}