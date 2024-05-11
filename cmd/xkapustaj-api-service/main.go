package main

import (
    "log"
    "os"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/jkloris/xkapustaj-webapi/api"
    "github.com/jkloris/xkapustaj-webapi/internal/xkapustaj_wl"
)

func main() {
    log.Printf("Server started")
    port := os.Getenv("XKAPUSTAJ_API_PORT")
    if port == "" {
        port = "8086"
    }
    environment := os.Getenv("XKAPUSTAJ_API_ENVIRONMENT")
    if !strings.EqualFold(environment, "production") { // case insensitive comparison
        gin.SetMode(gin.DebugMode)
    }
    engine := gin.New()
    engine.Use(gin.Recovery())
    // request routings
    xkapustaj_wl.AddRoutes(engine)
    engine.GET("/openapi", api.HandleOpenApi)
    engine.Run(":" + port)
}