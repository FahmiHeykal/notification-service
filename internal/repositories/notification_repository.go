package repositories

import (
    "gorm.io/gorm"
    "notification-service/internal/models"
)

type NotificationRepository struct {
    db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
    return &NotificationRepository{db: db}
}

func (r *NotificationRepository) Create(notification *models.Notification) error {
    return r.db.Create(notification).Error
}

func (r *NotificationRepository) FindByUserID(userID uint) ([]models.Notification, error) {
    var notifications []models.Notification
    err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&notifications).Error
    return notifications, err
}