package main

import (
	"log"
	"net/http"

	"github.com/ngamux/ngamux"
)

func main() {
	mux := ngamux.New()

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ngamux.Res(w).HTML("./index.html", ngamux.Map{
			"title":   "index",
			"message": "welcome to our site :)",
		})
	})

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
