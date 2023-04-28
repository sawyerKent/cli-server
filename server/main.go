package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	//"cli-server/server/handlers"
    "github.com/sawyerKent/cli-server/server/handlers"
)

func main() {
    // Initialize configuration using Viper
    viper.SetConfigFile(".env")
    viper.ReadInConfig()

    // Retrieve port number from configuration
    port := viper.GetInt("PORT")

    // Initialize Echo server
    e := echo.New()

    // Register routes
    e.GET("/", handlers.HiThere)
    e.GET("/heartbeat", handlers.Heartbeat)
    e.GET("/HappyLang", handlers.HappyLang)
    e.POST("/HappyLang", handlers.HappyLangPost)
    e.POST("/sendjson", handlers.ProcessJson)

    // Start server
    address := fmt.Sprintf(":%d", port)
    e.Logger.Fatal(e.Start(address))
}
