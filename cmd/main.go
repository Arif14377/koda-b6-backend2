package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

	// TODO: pindahkan pembacaan env ke preload satu file khusus.
	// TODO: Baru coba quick startup, buat sesuai kebutuhan.
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		DB:       0, //default DB
		Password: "",
	})
	defer rdb.Close()

	err = rdb.Set(ctx, "Key", "Value", 20*time.Second).Err()
	if err != nil {
		log.Fatal(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err == redis.Nil {
		fmt.Println("Key doesn't exists")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Key: ", val)
	}

	r := gin.Default()

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Ok",
		})
	})

	r.Run("localhost:8888")
}
