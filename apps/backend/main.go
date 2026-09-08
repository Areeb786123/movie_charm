package main

import (
	"backend/database"
	"backend/handler"
	"backend/kafka"
	"backend/repository"
	"backend/repository/outbox"
	"backend/routes"
	"backend/service"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	log.Println("🔥🔥🔥 MY NEW CODE IS RUNNING")

	db := database.ConnectDatabase()
	database.ConnectRedis()
	// 2. Repositories
	movieRepo := repository.CreateNewMovieRepo(db)

	// 3. Services
	movieService := service.CreateNewMovieService(movieRepo)
	commentService := service.CreateNewCommentService(movieRepo)
	ratingService := service.CreateRatingService(movieRepo)

	// 4. Handlers
	movieHandler := handler.NewMovieHandler(*movieService)
	commentHandler := handler.NewCommentHandler(*commentService)
	ratingHandler := handler.CreateRatingHandler(ratingService)

	brokers := strings.Split(getEnv("KAFKA_BROKERS", "localhost:9092"), ",")
	producer := kafka.NewKafkaProducer(brokers, "movie-rated")
	defer producer.Close()
	worker := kafka.NewOutboxWorker(outbox.CreateOutboxRepository(db), producer)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go runOutboxWorker(ctx, worker)

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
		ratingHandler,
	)

	server := &http.Server{Addr: ":8080", Handler: router}
	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("HTTP server shutdown failed: %v", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("HTTP server failed:", err)
	}
}

func runOutboxWorker(ctx context.Context, worker *kafka.OutboxWorker) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		if err := worker.Process(ctx); err != nil && !errors.Is(err, context.Canceled) {
			log.Printf("outbox worker error: %v", err)
		}

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}

func getEnv(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}
