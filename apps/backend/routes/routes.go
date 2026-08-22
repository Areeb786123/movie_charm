package routes

import (
	"backend/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	movieHandler *handler.MovieHandler,
	commentHandler *handler.CommentHandler,
) {

	// Movies
	router.POST("/movies", movieHandler.CreateMovie)
	router.GET("/movies", movieHandler.GetAllMovies)
	router.GET("/movies/:movieId", movieHandler.GetMovieById)
	router.DELETE("/movies/:movieId", movieHandler.DeleteMovie)

	// Comments
	router.POST("/movies/:movieId/comments", commentHandler.CreateComment)
	router.GET("/movies/:movieId/comments", commentHandler.GetAllComments)
}
