package routes

import (
    "notification-service/internal/controllers"
    "notification-service/internal/middleware"
    "notification-service/internal/websocket"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, authController *controllers.AuthController, userController *controllers.UserController, notificationController *controllers.NotificationController, hub *websocket.Hub, jwtSecret string) {
    api := router.Group("/api")
    {
        api.POST("/register", authController.Register)
        api.POST("/login", authController.Login)

        api.POST("/users", userController.CreateUser)

        authMiddleware := middleware.AuthMiddleware(jwtSecret)
        
        protected := api.Group("/")
        protected.Use(authMiddleware)
        {
            protected.GET("/ws", func(ctx *gin.Context) {
                userID := ctx.MustGet("userID").(uint)
                websocket.ServeWs(hub, ctx.Writer, ctx.Request, userID)
            })

            protected.GET("/notifications", notificationController.GetNotifications)
            protected.POST("/notifications/send", notificationController.SendNotification)
        }
    }
}