package main

import (
	"backend/database"
	"backend/handler"
	"backend/repository"
	"backend/routes"
	"backend/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectDatabase()

	// 2. Repositories
	movieRepo := repository.CreateNewMovieRepo(db)

	// 3. Services
	movieService := service.CreateNewMovieService(movieRepo)
	commentService := service.CreateNewCommentService(movieRepo)

	// 4. Handlers
	movieHandler := handler.NewMovieHandler(*movieService)
	commentHandler := handler.NewCommentHandler(*commentService)

	// Router
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.Status(204)
			c.Abort()
			return
		}
		c.Next()
	})

	routes.SetupRoutes(
		router,
		movieHandler,
		commentHandler,
	)

	// Start server
	router.Run(":8080")
}
