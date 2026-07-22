package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ngamux/ngamux"
)

func globalMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello from global middleware")
		next(w, r)
	}
}

func routeMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello from route middleware")
		next(w, r)
	}
}

func main() {
	mux := ngamux.New()
	mux.Use(globalMiddleware)

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello from root handler")
		ngamux.Res(w).
			Status(http.StatusOK).
			Text("GET /")
	})

	mux.Get("/users", ngamux.WithMiddlewares(routeMiddleware)(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello from users handler")
		ngamux.Res(w).
			Status(http.StatusOK).
			Text("GET /users")
	}))

	mux.With(routeMiddleware).Get("/products", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello from products handler")
		ngamux.Res(w).
			Status(http.StatusOK).
			Text("GET /products")
	})

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
