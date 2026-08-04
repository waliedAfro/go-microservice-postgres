package main

import (
	"fmt"
	"go-microservice-postgres/internal/config"
	"go-microservice-postgres/internal/database"
	"go-microservice-postgres/internal/handler"
	"go-microservice-postgres/internal/middleware"
	"go-microservice-postgres/internal/repository"
	"go-microservice-postgres/internal/service"
	"log"
	"net/http"
	"time"
)

func main() {

	cfg := config.Load()

	fmt.Printf("Database URL : %s", cfg.DatabaseURL)

	db := database.NewPostgres(cfg.DatabaseURL)
	defer db.Close()

	bookRepo := repository.NewPostgresBookRepository(db)

	bookService := service.NewBookService(bookRepo)

	bookHandler := handler.NewBookHandler(bookService)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /books", bookHandler.GetBooks)
	mux.HandleFunc("GET /books/{id}", bookHandler.GetBookByID)
	mux.HandleFunc("POST /books", bookHandler.CreateBook)
	mux.HandleFunc("DELETE /books/{id}", bookHandler.DeleteBook)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      middleware.Logging(middleware.CORS(mux)),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Server started on :8080")
	log.Fatal(server.ListenAndServe())

}
