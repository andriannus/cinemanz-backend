package routes

import (
	"cinemanz/controllers"

	"github.com/go-chi/chi"
)

// Routes create list API routes
func Routes() *chi.Mux {
	route := chi.NewRouter()

	route.Route("/api/v1", func(route chi.Router) {
		route.Route("/movies", func(route chi.Router) {
			route.Get("/", controllers.FetchMovies)
			route.Post("/", controllers.AddMovie)
			route.Get("/{movieID}", controllers.FetchMovie)
			route.Put("/{movieID}", controllers.UpdateMovie)
			route.Delete("/{movieID}", controllers.DeleteMovie)
		})
		route.Route("/theaters", func(route chi.Router) {
			route.Get("/", controllers.FetchTheaters)
			route.Post("/", controllers.AddTheater)
			route.Get("/{theaterID}", controllers.FetchTheater)
			route.Put("/{theaterID}", controllers.UpdateTheater)
			route.Delete("/{theaterID}", controllers.DeleteTheater)
		})
	})

	return route
}
