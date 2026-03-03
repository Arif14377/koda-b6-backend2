package service

import (
	"github.com/arif14377/koda-b6-backend2/internal/models"
	repository "github.com/arif14377/koda-b6-backend2/internal/repository/users"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetAll() []models.User {
	return *s.repo.GetAllUser()
}
