package handler

import (
	"net/http"
	"registration-smart-reward/internal/dto"
	"registration-smart-reward/internal/model"
	"registration-smart-reward/internal/service"

	"github.com/gin-gonic/gin"
)

type MemberHandler struct {
	svc service.MemberService
}

func NewMemberHandler(svc service.MemberService) *MemberHandler {
	return &MemberHandler{svc}
}

func (h *MemberHandler) Register(c *gin.Context) {
	var req dto.RegisterMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.Register(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registrasi member berhasil"})
}

func (h *MemberHandler) List(c *gin.Context) {
	var query dto.PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameter"})
		return
	}

	// Default values
	if query.Page < 1 {
		query.Page = 1
	}
	if query.Limit < 1 {
		query.Limit = 10
	}

	members, total, err := h.svc.ListPaginated(c.Request.Context(), query.Page, query.Limit, query.Search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data member"})
		return
	}

	if members == nil {
		members = []model.Member{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": members,
		"meta": gin.H{
			"total": total,
			"page":  query.Page,
			"limit": query.Limit,
		},
	})

}
