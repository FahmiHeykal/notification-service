package controllers

import (
    "net/http"
    "notification-service/internal/services"
    "github.com/gin-gonic/gin"
)

type UserController struct {
    userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
    return &UserController{userService: userService}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
    var request struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.userService.CreateUser(request.Username, request.Password); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"message": "user created"})
}