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

	mux.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "GET /users")
	})

	mux.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		c := ctx.New(w, r)
		c.Res().
			Status(http.StatusOK).
			Text("GET /products")
	})

	mux.Get("/api/status", func(w http.ResponseWriter, r *http.Request) {
		c := ctx.New(w, r)
		c.Res().
			Status(http.StatusOK).
			JSON(ngamux.Map{
				"status": "ok",
			})
	})

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
