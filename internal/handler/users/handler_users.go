package users

import (
	"github.com/arif14377/koda-b6-backend2/internal/service/users"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *users.Service
	*gin.Engine
}

func NewHandler(service *users.Service, api *gin.Engine) *Handler {
	return &Handler{
		service: service,
		Engine:  api,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("user")
	route.GET("/")
}
