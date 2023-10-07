package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/vargaadam23/rss-project-go/internal/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Main function is working")

	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		log.Fatal("No db url specified")
	}

	dbConnection, err := sql.Open("mysql", dbUrl)

	if err != nil {
		log.Println("Connection to the database failed %v", err)
	}

	db := database.New(dbConnection)

	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(
		db,
		5,
		time.Minute,
	)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	readyRouter := chi.NewRouter()
	readyRouter.Get("/ready", handleReadiness)
	readyRouter.Get("/error", handleError)

	readyRouter.Post("/users", apiCfg.handleCreateUser)
	readyRouter.Get("/users", apiCfg.middlewareAuth(apiCfg.handleGetUserByApiKey))

	readyRouter.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handleCreateFeed))
	readyRouter.Get("/feeds", apiCfg.middlewareAuth(apiCfg.handleGetFeeds))

	readyRouter.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handleCreateFeedFollow))
	readyRouter.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handleGetFeedFollows))
	readyRouter.Delete("/feed_follows/{feedFollowId}", apiCfg.middlewareAuth(apiCfg.handleDeleteFeedFollow))

	readyRouter.Get("/posts", apiCfg.middlewareAuth(apiCfg.handleGetUserPosts))

	router.Mount("/api/v1", readyRouter)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Println("Server started on port: ", port)
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Port is ", port)
}
