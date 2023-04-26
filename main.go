package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	h "students/pkg/http"
	"students/pkg/storage"
)

func main() {
	r := chi.NewRouter()

	strg := storage.NewStorage()
	c := h.Controller{Storage: strg}

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Post("/create", c.UserCreate)
	r.Post("/make_friends", c.UserAttach)
	r.Delete("/user", c.UserDelete)
	r.Get("/friends/{id}", c.UserGetFriends)
	r.Put("/{id}", c.UserUpdate)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
