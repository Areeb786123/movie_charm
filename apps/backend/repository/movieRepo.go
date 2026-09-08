package repository

import (
	"backend/dto"
	"backend/models"
	"context"
)

type MovieRepo interface {
	CreateMovie(movie *dto.CreateMovieRequest) error
	GetMoviesByID(movieId int) (models.Movies, error)
	GetMovies() ([]models.Movies, error)
	DeleteMovie(movieId int) error
	CreateComment(comment dto.CreateComments) error
	GetAllComment(movieId int, offset int, limit int) ([]models.Comment, error)
	RateMovie(ctx context.Context, data dto.CreateRating) error
}
