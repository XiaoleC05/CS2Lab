package handler

import (
	"net/http"
	"strconv"

	"github.com/XiaoleC05/CS2Lab/internal/db"
	"github.com/XiaoleC05/CS2Lab/internal/model"
	"github.com/gin-gonic/gin"
)

// MapHandler handles map-related HTTP requests
type MapHandler struct {
	repo *db.MapRepository
}

// NewMapHandler creates a new MapHandler
func NewMapHandler() *MapHandler {
	return &MapHandler{
		repo: db.NewMapRepository(),
	}
}

// GetAll handles GET /api/maps
func (h *MapHandler) GetAll(c *gin.Context) {
	maps, err := h.repo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, maps)
}

// GetByID handles GET /api/maps/:id
func (h *MapHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid map ID"})
		return
	}

	mapData, err := h.repo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "map not found"})
		return
	}

	c.JSON(http.StatusOK, mapData)
}

// LineupHandler handles lineup-related HTTP requests
type LineupHandler struct {
	repo *db.LineupRepository
}

// NewLineupHandler creates a new LineupHandler
func NewLineupHandler() *LineupHandler {
	return &LineupHandler{
		repo: db.NewLineupRepository(),
	}
}

// GetFiltered handles GET /api/lineups
func (h *LineupHandler) GetFiltered(c *gin.Context) {
	filter := model.LineupFilter{
		Limit:  50,
		Offset: 0,
	}

	if mapIDStr := c.Query("map_id"); mapIDStr != "" {
		if mapID, err := strconv.ParseInt(mapIDStr, 10, 64); err == nil {
			filter.MapID = &mapID
		}
	}

	if typeStr := c.Query("type"); typeStr != "" {
		filter.Type = &typeStr
	}

	if queryStr := c.Query("q"); queryStr != "" {
		filter.Query = &queryStr
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 {
			filter.Limit = limit
		}
	}

	if offsetStr := c.Query("offset"); offsetStr != "" {
		if offset, err := strconv.Atoi(offsetStr); err == nil && offset >= 0 {
			filter.Offset = offset
		}
	}

	lineups, err := h.repo.GetFiltered(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lineups)
}

// GetByID handles GET /api/lineups/:id
func (h *LineupHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lineup ID"})
		return
	}

	lineup, err := h.repo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "lineup not found"})
		return
	}

	c.JSON(http.StatusOK, lineup)
}

// FavoriteHandler handles favorite-related HTTP requests
type FavoriteHandler struct {
	repo *db.FavoriteRepository
}

// NewFavoriteHandler creates a new FavoriteHandler
func NewFavoriteHandler() *FavoriteHandler {
	return &FavoriteHandler{
		repo: db.NewFavoriteRepository(),
	}
}

// GetByUser handles GET /api/favorites
func (h *FavoriteHandler) GetByUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	favorites, err := h.repo.GetByUser(c.Request.Context(), userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, favorites)
}

// Add handles POST /api/favorites
func (h *FavoriteHandler) Add(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	var req struct {
		LineupID int64 `json:"lineup_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.repo.Add(c.Request.Context(), userID.(int64), req.LineupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "favorite added"})
}

// Remove handles DELETE /api/favorites/:lineupId
func (h *FavoriteHandler) Remove(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	lineupIDStr := c.Param("lineupId")
	lineupID, err := strconv.ParseInt(lineupIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lineup ID"})
		return
	}

	err = h.repo.Remove(c.Request.Context(), userID.(int64), lineupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "favorite removed"})
}

// NoteHandler handles note-related HTTP requests
type NoteHandler struct {
	repo *db.NoteRepository
}

// NewNoteHandler creates a new NoteHandler
func NewNoteHandler() *NoteHandler {
	return &NoteHandler{
		repo: db.NewNoteRepository(),
	}
}

// GetByLineup handles GET /api/notes/:lineupId
func (h *NoteHandler) GetByLineup(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	lineupIDStr := c.Param("lineupId")
	lineupID, err := strconv.ParseInt(lineupIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lineup ID"})
		return
	}

	note, err := h.repo.GetByLineup(c.Request.Context(), userID.(int64), lineupID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
		return
	}

	c.JSON(http.StatusOK, note)
}

// Upsert handles PUT /api/notes/:lineupId
func (h *NoteHandler) Upsert(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	lineupIDStr := c.Param("lineupId")
	lineupID, err := strconv.ParseInt(lineupIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lineup ID"})
		return
	}

	var req struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note, err := h.repo.Upsert(c.Request.Context(), userID.(int64), lineupID, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, note)
}

// HealthHandler handles health check requests
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
