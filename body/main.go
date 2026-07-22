package main

import (
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

	mux.Post("/", func(w http.ResponseWriter, r *http.Request) {
		c := ctx.New(w, r)

		in := make(map[string]string)
		if err := c.Req().JSON(&in); err != nil {
			c.Res().
				Status(http.StatusBadRequest).
				JSON(ngamux.Map{
					"error": err.Error(),
				})
			return
		}

		c.Res().Status(http.StatusCreated).JSON(in)
	})

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
