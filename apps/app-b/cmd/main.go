package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	_ "github.com/shsma/app-b/internal/database/postgres"
	"github.com/shsma/app-b/pkg/config"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	_, err := fmt.Fprintf(w, "Hello From APP-B!")
	if err != nil {
		http.Error(w, "Internal Server Problem.", http.StatusInternalServerError)
		return
	}
}

func main() {
	cfg := config.LoadServerConfig()
	cfg.LogServerConfig()

	http.HandleFunc("/hello", helloHandler)

	log.Info().Msgf("Starting app-b at port %s\n", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), nil); err != nil {
		log.Error().Err(err).Msg("Failed to start the server")
	}
}
