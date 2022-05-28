package main

import (
	"log"
	"net/http"
	"os"

	"github.com/KotaroYamazaki/go-event-notifier/handlers"
	"github.com/KotaroYamazaki/go-event-notifier/pkg/slack"
	"github.com/KotaroYamazaki/go-event-notifier/usecases"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	if os.Getenv("APP_ENV") == "local" {
		gin.SetMode(gin.DebugMode)
		router.Use(gin.Logger())
	}
	router.Use(gin.Recovery())

	slackClient := slack.New(os.Getenv("SLACK_WEBHOOK_URL"))
	slackUsecase := usecases.NewSlackUsecase(slackClient)
	slackHandler := handlers.NewNotifierHandler(slackUsecase)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})
	webhook := router.Group("slack")
	{
		webhook.POST("/error", slackHandler.Notify)
	}
	router.Run(":" + os.Getenv("PORT"))
}
