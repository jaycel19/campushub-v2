package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{s}
}

type CreatePostRequest struct {
	Content string `json:"content" binding:"required"`
}

func (h *Handler) GetFeed(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	posts, err := h.service.GetFeed(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":  page,
		"limit": limit,
		"data":  posts,
	})
}

func (h *Handler) CreatePost(c *gin.Context) {
	var req CreatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userIDStr := userID.(string)

	parsedID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id",
		})
		return
	}

	post := Post{
		UserID:  parsedID,
		Content: req.Content,
	}

	err = h.service.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}
