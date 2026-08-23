package repository

import (
	"backend/dto"
	"backend/models"
)

type MovieRepo interface {
	CreateMovie(movie *dto.CreateMovieRequest) error
	GetMoviesByID(movieId int) (models.Movies, error)
	GetMovies() ([]models.Movies, error)
	DeleteMovie(movieId int) error
	CreateComment(comment dto.CreateComments) error
	GetAllComment(movieId int, offset int, limit int) ([]models.Comment, error)
}
