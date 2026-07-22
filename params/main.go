package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ngamux/ctx"
	"github.com/ngamux/ngamux"
)

func main() {
	mux := ngamux.New()

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ngamux.Res(w).
			Status(http.StatusOK).
			Text("GET /")
	})

	users := mux.Group("/users")
	users.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		ngamux.Res(w).
			Status(http.StatusOK).
			Text(fmt.Sprintf("GET /users/{id} with id: %s", r.PathValue("id")))
	})

	users.Get("/{id}/{slug}", func(w http.ResponseWriter, r *http.Request) {
		c := ctx.New(w, r)
		c.Res().
			Status(http.StatusOK).
			JSON(ngamux.Map{
				"id":   c.Req().PathValue("id"),
				"slug": c.Req().PathValue("slug"),
			})

	})

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
