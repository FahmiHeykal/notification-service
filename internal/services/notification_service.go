package services

import (
    "notification-service/internal/models"
    "notification-service/internal/repositories"
    "notification-service/internal/websocket"
)

type NotificationService struct {
    notificationRepo *repositories.NotificationRepository
    hub             *websocket.Hub
}

func NewNotificationService(notificationRepo *repositories.NotificationRepository, hub *websocket.Hub) *NotificationService {
    return &NotificationService{
        notificationRepo: notificationRepo,
        hub:             hub,
    }
}

func (s *NotificationService) SendNotification(userID uint, message string) error {
    notification := &models.Notification{
        UserID:  userID,
        Message: message,
        Read:    false,
    }

    if err := s.notificationRepo.Create(notification); err != nil {
        return err
    }

    s.hub.SendToUser(userID, message)
    return nil
}

func (s *NotificationService) GetUserNotifications(userID uint) ([]models.Notification, error) {
    return s.notificationRepo.FindByUserID(userID)
}