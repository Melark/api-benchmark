package main

import (
	"benchmark/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func main() {

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data := models.Person{Name: "John", Surname: "Doe"}

		render.JSON(w, r, data)
	})
	http.ListenAndServe(":3000", r)
}
