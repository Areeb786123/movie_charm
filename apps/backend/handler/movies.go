package handler

import (
	"backend/dto"
	"backend/service"
	"backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	s service.MovieService
}

func NewMovieHandler(s service.MovieService) *MovieHandler {
	return &MovieHandler{
		s: s,
	}
}

func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var req dto.CreateMovieRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.s.CreateMovie(req)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Movie created successfully",
	})
}

func (h *MovieHandler) GetMovieById(c *gin.Context) {
	movieID, err := strconv.Atoi(c.Param("movieId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, err := h.s.GetMovieById(movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *MovieHandler) GetAllMovies(c *gin.Context) {
	movies, err := h.s.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	movieID, err := strconv.Atoi(c.Param("movieId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.s.DeleteMovie(movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Movie deleted successfully",
	})
}
