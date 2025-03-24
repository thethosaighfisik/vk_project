package main 

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"vk_project/handlers"
)

func main(){
	r := chi.NewRouter()

	r.Post("/kv", handlers.HandlerPost)
	r.Put("/kv/{id}", handlers.HandlerPut)
	r.Get("/kv/{id}", handlers.HandlerGet)
	r.Delete("/kv/{id}", handlers.HandlerDelete)

	http.ListenAndServe(":8080", r)

}
