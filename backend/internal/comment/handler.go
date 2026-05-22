package comment

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid post id",
		})
		return
	}

	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userIDStr, ok := userIDValue.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id type",
		})
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id",
		})
		return
	}

	comment := Comment{
		PostID:  postID,
		UserID:  userID,
		Content: req.Content,
	}

	err = h.service.CreateComment(&comment)
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

	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid post id",
		})
		return
	}

	comments, err := h.service.GetComments(postID)
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

	commentID, err := uuid.Parse(commentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid comment id",
		})
		return
	}

	userIDValue, _ := c.Get("user_id")
	userIDStr := userIDValue.(string)

	userID, _ := uuid.Parse(userIDStr)

	err = h.service.DeleteComment(commentID, userID)
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
