package handler

import (
	"github.com/aimericsr/social-network-auth-api/api/controller"
	"github.com/aimericsr/social-network-auth-api/api/middleware"
	"github.com/aimericsr/social-network-auth-api/token"
	"github.com/go-chi/chi"
)

func NewHandler(TokenMaker token.Maker) *chi.Mux {
	router := chi.NewRouter()

	router.MethodNotAllowed(controller.MethodNotAllowedHandler)
	router.NotFound(controller.NotFoundHandler)

	router.Use(middleware.SimpleLogger)
	//router.Use(middleware.BasicAuth)

	router.Route("/v1", func(router chi.Router) {
		router.Get("/health", controller.HealthCheck)

		router.Route("/users", func(router chi.Router) {
			router.Get("/", controller.GetArtists)

			router.Route("/users", func(router chi.Router) {
				router.Get("/", controller.GetArtists)
			})
		})
	})

	return router
}
