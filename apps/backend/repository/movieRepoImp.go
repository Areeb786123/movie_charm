package repository

import (
	"backend/dto"
	"backend/models"

	"gorm.io/gorm"
)

type movieRepo struct {
	db *gorm.DB
}

func CreateNewMovieRepo(db *gorm.DB) MovieRepo {
	return &movieRepo{
		db: db,
	}
}

func (m *movieRepo) CreateMovie(movie *dto.CreateMovieRequest) error {
	newMovie := models.Movies{
		MovieName:  movie.MovieName,
		UploadedBy: movie.UploadedBy,
	}

	result := m.db.Create(&newMovie)
	return result.Error
}

func (m *movieRepo) GetMoviesByID(movieId int) (models.Movies, error) {
	var movie models.Movies
	result := m.db.First(&movie, movieId)

	if result.Error != nil {
		return movie, result.Error
	}
	return movie, nil
}

func (m *movieRepo) GetMovies() ([]models.Movies, error) {
	var movie []models.Movies

	result := m.db.Find(&movie)
	if result.Error != nil {
		return nil, result.Error
	}

	return movie, nil
}

func (m *movieRepo) DeleteMovie(movieId int) error {
	var movie models.Movies
	result := m.db.Delete(&movie, movieId)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *movieRepo) CreateComment(comment dto.CreateComments) error {
	newComment := models.Comments{
		MovieId: comment.MovieID,
		Comment: comment.Comment,
	}
	result := m.db.Create(&newComment)
	return result.Error
}

func (m *movieRepo) GetAllComment(movieId int, offset int, limit int) ([]models.Comment, error) {
	var comments []models.Comment
	result := m.db.Where("movie_id= ?", movieId).Offset(offset).Limit(limit).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}
