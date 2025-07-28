package models

import "gorm.io/gorm"

type Notification struct {
    gorm.Model
    UserID  uint
    User    User   `gorm:"foreignKey:UserID"`
    Message string
    Read    bool
}