package router

import "github.com/go-chi/chi"

func Init() *chi.Mux {
	var router = chi.NewRouter()
	return router
}
