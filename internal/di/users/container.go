package di

import (
	repository "github.com/arif14377/koda-b6-backend2/internal/repository/users"
	service "github.com/arif14377/koda-b6-backend2/internal/service/users"
)

type Container struct {
	userRepo    *repository.Repository
	userService *service.Service
}
