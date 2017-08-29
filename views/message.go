package views

import (
	"errors"
	"fmt"
	"net/http"

	"../models"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/go-resty/resty"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func PushNotifications(c *gin.Context) {
	var message models.Message
	err := c.BindJSON(&message)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("Bad Json"))
		return
	}
	log.Info(fmt.Sprintf("%v", message))

	url := viper.GetString("fcm.hot")
	fmt.Println("URL:>", url)

	jsonString, err := json.Marshal(message)

	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "key="+viper.GetString("fcm.keyServer")).
		SetBody(jsonString).
		Post(url)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// explore response object
	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Recevied At: %v", resp.ReceivedAt())
	fmt.Printf("\nResponse Body: %v", resp)

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
