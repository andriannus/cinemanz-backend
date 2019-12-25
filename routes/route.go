package routes

import (
	"github.com/go-chi/chi"

	"cinemanz/controllers"
	"cinemanz/middlewares"
)

// Routes create list API routes
func Routes() *chi.Mux {
	route := chi.NewRouter()

	route.Use(middlewares.Cors().Handler)

	route.Route("/v1", func(route chi.Router) {
		route.Route("/movies", func(route chi.Router) {
			route.Get("/", controllers.FetchMovies)
			route.Get("/{movieID}", controllers.FetchMovie)

			route.With(middlewares.IsAuthenticated).Group(func(route chi.Router) {
				route.Post("/", controllers.AddMovie)
				route.Put("/{movieID}", controllers.UpdateMovie)
				route.Delete("/{movieID}", controllers.DeleteMovie)
			})
		})

		route.Route("/theaters", func(route chi.Router) {
			route.Get("/", controllers.FetchTheaters)
			route.Get("/{theaterID}", controllers.FetchTheater)

			route.With(middlewares.IsAuthenticated).Group(func(route chi.Router) {
				route.Post("/", controllers.AddTheater)
				route.Put("/{theaterID}", controllers.UpdateTheater)
				route.Delete("/{theaterID}", controllers.DeleteTheater)
			})
		})

		route.Route("/user", func(route chi.Router) {
			route.Post("/login", controllers.Login)
			route.Post("/register", controllers.Register)
		})
	})

	return route
}
