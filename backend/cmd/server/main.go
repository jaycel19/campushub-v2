package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jaycel19/campushub/backend/internal/post"
	"github.com/jaycel19/campushub/backend/pkg/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loding .env file")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	postRepo := post.NewRepository(db)
	postService := post.NewService(postRepo)
	postHandler := post.NewHandler(postService)

	r := gin.Default()

	r.GET("/feed", postHandler.GetFeed)
	r.POST("/posts", postHandler.CreatePost)

	port := os.Getenv("HTTP_PORT")
	r.Run(":" + port)
}
