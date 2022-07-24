package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RafaelMaizonave/api-template-base/configs"
	"github.com/RafaelMaizonave/api-template-base/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	router := chi.NewRouter()

	router.Post("/", handlers.Create)
	router.Put("/{id}", handlers.Update)
	router.Delete("/{id}", handlers.Delete)
	router.Get("/", handlers.List)
	router.Get("/{id}", handlers.GetById)

	for {
		log.Println(http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router))
	}

}
