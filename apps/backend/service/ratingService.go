package service

import (
	"backend/dto"
	"backend/repository"
	"context"
)

type RatingService struct {
	repo repository.MovieRepo
}

func CreateRatingService(repo repository.MovieRepo) *RatingService {
	return &RatingService{repo: repo}
}

func (r *RatingService) RateMovie(ctx context.Context, data dto.CreateRating) error {
	return r.repo.RateMovie(ctx, data)
}
