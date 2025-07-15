package handler

import (
	"net/http"
	"registration-smart-reward/internal/repo"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
    memberRepo repo.MemberRepository
    paketRepo  repo.PaketRepository
}


func NewDashboardHandler(mRepo repo.MemberRepository, pRepo repo.PaketRepository) *DashboardHandler {
    return &DashboardHandler{
        memberRepo: mRepo,
        paketRepo:  pRepo,
    }
}


func (h *DashboardHandler) Overview(c *gin.Context) {
	ctx := c.Request.Context()

	memberCount, err := h.memberRepo.CountAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal hitung member"})
		return
	}

	paketCount, err := h.paketRepo.CountAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal hitung paket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_member": memberCount,
		"total_paket":  paketCount,
	})
}
