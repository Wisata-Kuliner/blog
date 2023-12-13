package main

import (
	"blog/config"
	"blog/handler"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}
	h := handler.New(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/posts", h.BlogPostList)
	r.Post("/posts", h.BlogPostCreate)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("Server started at port " + port)
	http.ListenAndServe(":"+port, r)
}
