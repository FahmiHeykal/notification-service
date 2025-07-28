package controllers

import (
    "net/http"
    "notification-service/internal/services"
    "github.com/gin-gonic/gin"
)

type NotificationController struct {
    notificationService *services.NotificationService
}

func NewNotificationController(notificationService *services.NotificationService) *NotificationController {
    return &NotificationController{notificationService: notificationService}
}

func (c *NotificationController) SendNotification(ctx *gin.Context) {
    var request struct {
        UserID  uint   `json:"user_id"`
        Message string `json:"message"`
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.notificationService.SendNotification(request.UserID, request.Message); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "notification sent"})
}

func (c *NotificationController) GetNotifications(ctx *gin.Context) {
    userID := ctx.MustGet("userID").(uint)

    notifications, err := c.notificationService.GetUserNotifications(userID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"notifications": notifications})
}