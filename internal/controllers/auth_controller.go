package controllers

import (
    "net/http"
    "notification-service/internal/services"
    "github.com/gin-gonic/gin"
)

type AuthController struct {
    authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
    return &AuthController{authService: authService}
}

func (c *AuthController) Register(ctx *gin.Context) {
    var request struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.authService.Register(request.Username, request.Password); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func (c *AuthController) Login(ctx *gin.Context) {
    var request struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := c.authService.Login(request.Username, request.Password)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"token": token})
}