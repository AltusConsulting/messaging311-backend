package main

import (
	"fmt"
	"net/http"

	"./views"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	initializeApp()
	router := getEngine()
	router.Run(fmt.Sprintf(":%s", viper.GetString("server.port")))
}
func getEngine() *gin.Engine {
	router := gin.Default()
	router.Use(gin.ErrorLoggerT(gin.ErrorTypePrivate))
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := router.Group("api/messaging/v1")
	{
		v1.POST("/notifications", views.PushNotifications)
	}

	return router
}

func initializeApp() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
