package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "List Users:",
		"results": h.service.GetAll(),
	})
}
