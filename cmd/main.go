package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// TODO: buat handle untuk redis

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load godotenv: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(os.Getenv("REDIS_HOST"), ":", os.Getenv("REDIS_HOST")),
	})
	defer rdb.Close()

	r := gin.Default()

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Ok",
		})
	})

	r.Run("localhost:8888")
}
