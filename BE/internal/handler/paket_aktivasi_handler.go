package handler

import (
	"net/http"
	"registration-smart-reward/internal/dto"
	"registration-smart-reward/internal/service"

	"github.com/gin-gonic/gin"
)

type PaketAktivasiHandler struct {
	svc service.PaketAktivasiService
}

func NewPaketAktivasiHandler(svc service.PaketAktivasiService) *PaketAktivasiHandler {
	return &PaketAktivasiHandler{svc}
}

func (h *PaketAktivasiHandler) Activate(c *gin.Context) {
	var req dto.PaketAktivasiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.Activate(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "paket berhasil diaktivasi"})
}
