package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jaycel19/campushub/backend/internal/auth"
	"github.com/jaycel19/campushub/backend/internal/post"
	"github.com/jaycel19/campushub/backend/pkg/database"
	"github.com/jaycel19/campushub/backend/pkg/middlewares"
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

	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	postRepo := post.NewRepository(db)
	postService := post.NewService(postRepo)
	postHandler := post.NewHandler(postService)

	r := gin.Default()

	// post routes
	r.GET("/feed", postHandler.GetFeed)
	r.POST("/posts", middlewares.AuthMiddleware(), postHandler.CreatePost) // protected route

	// auth routes
	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)
	r.GET("/users", middlewares.AuthMiddleware(), authHandler.GetAll) // protected route
	r.GET("/me", middlewares.AuthMiddleware(), authHandler.GetMe)     // protected route

	port := os.Getenv("HTTP_PORT")
	r.Run(":" + port)
}
