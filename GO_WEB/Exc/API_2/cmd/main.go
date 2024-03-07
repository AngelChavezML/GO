package main

import (
	"fmt"
	"go_web/post/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	handler := handlers.DefaultTask{}
	router := chi.NewRouter()
	router.Post("/greeting", handler.Create())
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
	}

}
