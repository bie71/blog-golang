package app

import (
	"blog/config"
	"net/http"

	"github.com/rs/zerolog/log"
)

func RunServer() error {
	cfg := config.NewConfig()

	_, err := cfg.ConectionPostgres()

	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to database")
		return err
	}

	// Set up your server (e.g., routes, middleware)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Start the server
	log.Info().Msgf("Starting server on port %s", cfg.APP.APP_PORT)
	log.Info().Msgf("Environment: %s", cfg.APP.APP_ENV)

	return http.ListenAndServe(":8080", nil)
}
