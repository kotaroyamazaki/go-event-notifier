package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/KotaroYamazaki/go-event-notifier/entites"
	"github.com/gin-gonic/gin"
)

type NotifierHandler struct {
	notifier entites.Notifier
}

func NewNotifierHandler(notifier entites.Notifier) *NotifierHandler {
	return &NotifierHandler{
		notifier: notifier,
	}
}

func (h *NotifierHandler) Notify(c *gin.Context) {
	var appLog *entites.AppLog
	var psm *entites.PubSubMessage
	if err := c.ShouldBindJSON(psm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := json.Unmarshal(psm.Message.Data, appLog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.notifier.Notify(c, appLog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}
