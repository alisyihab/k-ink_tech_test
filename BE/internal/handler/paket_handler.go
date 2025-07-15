package handler

import (
	"net/http"
	"registration-smart-reward/internal/dto"
	"registration-smart-reward/internal/service"

	"github.com/gin-gonic/gin"
)

type PaketHandler struct {
	svc service.PaketService
}

func NewPaketHandler(svc service.PaketService) *PaketHandler {
	return &PaketHandler{svc}
}

func (h *PaketHandler) Create(c *gin.Context) {
	var req dto.CreatePaketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.Create(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "paket berhasil ditambahkan"})
}

func (h *PaketHandler) ListActive(c *gin.Context) {
	pakets, err := h.svc.ListActive(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pakets})
}
