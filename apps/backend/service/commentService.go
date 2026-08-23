package service

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
)

type CommentService struct {
	repo repository.MovieRepo
}

func CreateNewCommentService(repository repository.MovieRepo) *CommentService {
	return &CommentService{
		repo: repository,
	}
}

func (r *CommentService) AddCommentOnMovie(comment dto.CreateComments) (string, error) {
	err := r.repo.CreateComment(comment)
	if err != nil {
		return "some error occur", err
	}
	//here we will call our ai model so our ai model will reply according the comment setiments
	return "success", nil
}

func (r *CommentService) GetAllComment(movieId int, offset int, limit int) ([]models.Comment, error) {
	result, err := r.repo.GetAllComment(movieId, offset, limit)
	if err != nil {
		return nil, err
	}
	return result, nil
}
