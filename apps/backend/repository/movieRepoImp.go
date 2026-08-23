package repository

import (
	"backend/dto"
	"backend/models"
	"fmt"

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
