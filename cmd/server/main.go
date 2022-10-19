package main

import (
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	log.Println("starting")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		//write response
	})

	r.Get("/ingredients", func(w http.ResponseWriter, r *http.Request) {
		//json.NewEncoder(w).Encode(...)
		//write response
	})

	r.Post("/order-buns", func(w http.ResponseWriter, r *http.Request) {
		//json.NewDecoder(...).Decode(...)
		//write response
	})

	r.Post("/slice", func(w http.ResponseWriter, r *http.Request) {
		// range
		//write response
	})

	r.Post("/slice-faster", func(w http.ResponseWriter, r *http.Request) {
		//var wg sync.WaitGroup
		//wg.Add(1)
		//wg.Done()
		//wg.Wait()
		//write response
	})

	r.Post("/mix", func(w http.ResponseWriter, r *http.Request) {
		//write response
	})

	r.Post("/cook", func(w http.ResponseWriter, r *http.Request) {
		//write response
	})

	r.Post("/assemble", func(w http.ResponseWriter, r *http.Request) {
		//b := burger.New(burger.Ingredients())
		//w.Write([]byte(b.Assemble()))
		//write response
	})

	http.ListenAndServe(":8081", r) //what if port is used, how we should handle it?
}
