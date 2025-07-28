package main

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "notification-service/internal/config"
    "notification-service/internal/controllers"
    "notification-service/internal/models"
    "notification-service/internal/repositories"
    "notification-service/internal/routes"
    "notification-service/internal/services"
    "notification-service/internal/websocket"
    "github.com/gin-gonic/gin"
)

func main() {
    cfg := config.LoadConfig()

    db, err := gorm.Open(postgres.Open(
        "host="+cfg.DBHost+
            " port="+cfg.DBPort+
            " user="+cfg.DBUser+
            " password="+cfg.DBPassword+
            " dbname="+cfg.DBName+
            " sslmode=disable"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&models.User{}, &models.Notification{})

    hub := websocket.NewHub()
    go hub.Run()

    userRepo := repositories.NewUserRepository(db)
    notificationRepo := repositories.NewNotificationRepository(db)

    authService := services.NewAuthService(userRepo, cfg.JWTSecret)
    userService := services.NewUserService(userRepo)
    notificationService := services.NewNotificationService(notificationRepo, hub)

    authController := controllers.NewAuthController(authService)
    userController := controllers.NewUserController(userService)
    notificationController := controllers.NewNotificationController(notificationService)

    router := gin.Default()
    routes.SetupRoutes(router, authController, userController, notificationController, hub, cfg.JWTSecret)

    router.Run(":8080")
}