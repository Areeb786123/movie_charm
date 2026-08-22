package service

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	"errors"
)

type MovieService struct {
	repo repository.MovieRepo
}

func CreateNewMovieService(repo repository.MovieRepo) *MovieService {
	return &MovieService{
		repo: repo,
	}
}
func (r *MovieService) CreateMovie(movie dto.CreateMovieRequest) error {
	if movie.MovieName == "" || movie.UploadedBy == "" {
		return errors.New("movie name or movie uploaded_by is required")
	}

	result := r.repo.CreateMovie(&movie)

	if result != nil {
		return errors.New(result.Error())
	}

	return nil
}

func (r *MovieService) GetAllMovies() ([]models.Movies, error) {
	result, err := r.repo.GetMovies()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *MovieService) GetMovieById(movieId int) (models.Movies, error) {
	result, err := r.repo.GetMoviesByID(movieId)
	if err != nil {
		return models.Movies{}, err
	}

	return result, nil
}

func (r *MovieService) DeleteMovie(movieId int) error {
	err := r.repo.DeleteMovie(movieId)
	if err != nil {
		return err
	}
	return nil
}
