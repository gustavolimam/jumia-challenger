package web

import (
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
	"github.com/jumia-challenger/internal/environment"

	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
)

func Run() {
	log.Info().Msg("Starting web application...")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:" + os.Getenv(environment.WebPort), "http://localhost:3000"}, // All origins
		AllowedMethods: []string{"*"},                                                                           // Allowing only get, just an example
		AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	// Criação da variavel de rotas HTTP
	router := mux.NewRouter()
	if router == nil {
		log.Error().Msg("Error trying to create a new Mux Router")
		return
	}

	// Carrega os arquivos estáticos do Front
	fs := http.FileServer(http.Dir(path.Join("web/frontend/build")))
	router.PathPrefix("/").Handler(fs)
	if fs == nil {
		log.Error().Msg("Error trying to load static files of web application")
	}

	log.Info().Msgf("Server running in port: %v", os.Getenv(environment.WebPort))
	if err := http.ListenAndServe(":"+os.Getenv(environment.WebPort), c.Handler(router)); err != nil {
		log.Error().Msg("Error trying to start the web application server")
	}
}
