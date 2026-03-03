package users

import (
	"github.com/arif14377/koda-b6-backend2/internal/models"
	"github.com/arif14377/koda-b6-backend2/internal/repository/users"
)

type Service struct {
	repo *users.Repository
}

func NewService(repo *users.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetAll() []models.User {
	return *s.repo.GetAllUser()
}
