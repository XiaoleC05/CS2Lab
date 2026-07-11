package handler

import (
	"net/http"
	"strconv"

	"github.com/XiaoleC05/CS2Lab/internal/model"
	"github.com/gin-gonic/gin"
)

// CreateMap POST /api/admin/maps
func (h *MapHandler) CreateMap(c *gin.Context) {
	var req struct {
		Name         string `json:"name" binding:"required"`
		DisplayName  string `json:"display_name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m, err := h.repo.Create(c.Request.Context(), req.Name, req.DisplayName)
	if err != nil {
		respondInternalError(c, err)
		return
	}

	c.JSON(http.StatusCreated, m)
}

// DeleteMap DELETE /api/admin/maps/:id
func (h *MapHandler) DeleteMap(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid map ID"})
		return
	}

	if err := h.repo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "map not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "map deleted"})
}

// CreateLineup POST /api/admin/lineups
func (h *LineupHandler) CreateLineup(c *gin.Context) {
	var req model.Lineup
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lineup, err := h.repo.Create(c.Request.Context(), &req)
	if err != nil {
		respondInternalError(c, err)
		return
	}

	c.JSON(http.StatusCreated, lineup)
}