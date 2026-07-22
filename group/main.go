package main

import (
	"log"
	"net/http"

	"github.com/ngamux/ngamux"
)

func main() {
	mux := ngamux.New()

	users := mux.Group("/users")
	users.Get("", func(w http.ResponseWriter, r *http.Request) {
		ngamux.Res(w).
			Status(http.StatusOK).
			Text("GET /users")
	})

	admins := users.Group("/admins")
	admins.Get("", func(w http.ResponseWriter, r *http.Request) {
		ngamux.Res(w).
			Status(http.StatusOK).
			Text("GET /users/admins")
	})

	log.Println("Server running :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
