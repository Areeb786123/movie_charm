package repository

import (
	"backend/dto"
	"backend/models"
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MovieRep struct {
	Db *gorm.DB
}

func CreateNewMovieRepo(db *gorm.DB) MovieRepo {
	return &MovieRep{
		Db: db,
	}
}

func (m *MovieRep) CreateMovie(movie *dto.CreateMovieRequest) error {
	fmt.Println(movie)
	newMovie := models.Movies{
		MovieName:  movie.MovieName,
		UploadedBy: movie.UploadedBy,
		ImageUrl:   movie.ImageUrl,
		MovieLink:  movie.MovieLink,
		TrailorUrl: movie.TrailorUrl,
	}

	result := m.Db.Create(&newMovie)
	return result.Error
}

func (m *MovieRep) GetMoviesByID(movieId int) (models.Movies, error) {
	var movie models.Movies
	result := m.Db.First(&movie, movieId)

	if result.Error != nil {
		return movie, result.Error
	}
	return movie, nil
}

func (m *MovieRep) GetMovies() ([]models.Movies, error) {
	var movie []models.Movies

	result := m.Db.Find(&movie)
	if result.Error != nil {
		return nil, result.Error
	}

	return movie, nil
}

func (m *MovieRep) DeleteMovie(movieId int) error {
	var movie models.Movies
	result := m.Db.Delete(&movie, movieId)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MovieRep) CreateComment(comment dto.CreateComments) error {
	newComment := models.Comments{
		MovieId: comment.MovieID,
		Comment: comment.Comment,
	}
	result := m.Db.Create(&newComment)
	return result.Error
}

func (m *MovieRep) GetAllComment(movieId int, offset int, limit int) ([]models.Comment, error) {
	var comments []models.Comment
	result := m.Db.Where("movie_id= ?", movieId).Offset(offset).Limit(limit).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

func (m *MovieRep) RateMovie(
	ctx context.Context,
	data dto.CreateRating,
) error {

	if data.Rating < 1 || data.Rating > 5 {
		return fmt.Errorf("rating must be between 1 and 5")
	}

	return m.Db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		// 1. Save rating
		rating := models.Rating{
			MovieId: data.MovieID,
			Rating:  data.Rating,
		}

		if err := tx.Create(&rating).Error; err != nil {
			return fmt.Errorf("create rating: %w", err)
		}

		// 2. Recalculate movie rating aggregate
		var result struct {
			Average float64
			Count   int64
		}

		if err := tx.Model(&models.Rating{}).
			Where("movie_id = ?", data.MovieID).
			Select("AVG(rating) AS average, COUNT(*) AS count").
			Scan(&result).Error; err != nil {
			return fmt.Errorf("calculate rating aggregate: %w", err)
		}

		// 3. Update movie aggregate
		if err := tx.Model(&models.Movies{}).
			Where("movie_id = ?", data.MovieID).
			Updates(map[string]interface{}{
				"average_rating": result.Average,
				"rating_count":   result.Count,
			}).Error; err != nil {
			return fmt.Errorf("update movie rating aggregate: %w", err)
		}

		// 4. Create movie-rated event
		event := models.MovieRatedEvent{
			EventID: uuid.New().String(),
			MovieID: data.MovieID,
			Rating:  data.Rating,
		}

		payload, err := json.Marshal(event)
		if err != nil {
			return fmt.Errorf("marshal movie rated event: %w", err)
		}

		// 5. Save event in outbox using the SAME transaction
		outboxEvent := models.OutboxEvent{
			EventType:   "movie.rated",
			AggregateID: int64(data.MovieID),
			Payload:     payload,
		}

		if err := tx.Create(&outboxEvent).Error; err != nil {
			return fmt.Errorf("create outbox event: %w", err)
		}

		return nil
	})
}
