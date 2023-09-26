package learning_go

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"learning-go/configs"
	"learning-go/handlers"
	"net/http"
)

func main(){

	err := configs.Load()
	if err != nil{
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}