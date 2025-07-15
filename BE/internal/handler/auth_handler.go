package handler

import (
	"net/http"
	"registration-smart-reward/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{ svc service.UserService }

func NewAuthHandler(svc service.UserService) *AuthHandler { return &AuthHandler{svc} }

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.svc.Login(c, req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
