package comment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

type CreateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

func NewHandler(s Service) *Handler {
	return &Handler{s}
}

func (h *Handler) CreateComment(c *gin.Context) {
	var req CreateCommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	postIDStr := c.Param("id")
	userIDStr := c.GetString("user_id")

	comment, err := h.service.CreateComment(
		postIDStr,
		userIDStr,
		req.Content,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "comment created",
		"data":    comment,
	})
}

func (h *Handler) GetComments(c *gin.Context) {
	postIDStr := c.Param("id")

	comments, err := h.service.GetComments(postIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comments,
	})
}

func (h *Handler) DeleteComment(c *gin.Context) {
	commentIDStr := c.Param("id")
	userIDStr := c.GetString("user_id")
	err := h.service.DeleteComment(commentIDStr, userIDStr)
	if err != nil {
		if err == ErrUnauthorized {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "forbidden",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "comment deleted",
	})
}
