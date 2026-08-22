package handler

import (
	"backend/dto"
	"backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	s service.CommentService
}

func NewCommentHandler(s service.CommentService) *CommentHandler {
	return &CommentHandler{
		s: s,
	}
}

// POST /movies/:movieId/comments
func (h *CommentHandler) CreateComment(c *gin.Context) {

    movieID, err := strconv.Atoi(c.Param("movieId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "invalid movie id",
        })
        return
    }

    var req dto.CreateComments

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    req.MovieID = movieID

    comment, err := h.s.AddCommentOnMovie(req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, comment)
}

// GET /movies/:movieId/comments?offset=0&limit=10
func (h *CommentHandler) GetAllComments(c *gin.Context) {

	movieID, err := strconv.Atoi(c.Param("movieId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid movie id",
		})
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid offset",
		})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid limit",
		})
		return
	}

	comments, err := h.s.GetAllComment(movieID, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, comments)
}
