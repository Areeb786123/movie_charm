package service

import (
	"backend/database"
	"backend/dto"
	"backend/models"
	"backend/repository"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
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
	cacheKey := models.GET_ALL_MOVIES
	ctx := context.Background()
	
	cacheData, err := database.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var movies []models.Movies
		err := json.Unmarshal([]byte(cacheData), &movies)
		if err != nil {
			return nil, err
		}
		log.Println("Data coming from redis")
		return movies, nil
	}

	//redis fail now get data from data it means data does not exis in redis
	result, err := r.repo.GetMovies()
	fmt.Println("Data coming from DB")
	if err != nil {
		return nil, err
	}
	//also store result in redis
	data, err := json.Marshal(result)
	if err == nil {
		err := database.RedisClient.Set(
			ctx,
			cacheKey,
			data,
			5*time.Minute,
		).Err()
		log.Println("data saved iun redis successfully")
		if err != nil {
			log.Println("some error occur saving data into redis ", err)
		}
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
