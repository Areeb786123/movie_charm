package handler

import (
	"backend/dto"
	"backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RatingHandler struct {
	service *service.RatingService
}

func CreateRatingHandler(service *service.RatingService) *RatingHandler {
	return &RatingHandler{service: service}
}

// RateMovie handles POST /movies/:movieId/rating.
func (h *RatingHandler) RateMovie(c *gin.Context) {
	movieID, err := strconv.Atoi(c.Param("movieId"))
	if err != nil || movieID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid movie id"})
		return
	}

	var req dto.CreateRating
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Rating < 1 || req.Rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rating must be between 1 and 5"})
		return
	}
	req.MovieID = movieID

	if err := h.service.RateMovie(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "movie rated successfully"})
}
