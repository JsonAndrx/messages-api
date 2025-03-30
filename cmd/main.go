package main

import (
	"api-messages/api/routes"

	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	routes.SetupRoutes()
	errServ := http.ListenAndServe(":8080", nil)
	if errServ != nil {
		log.Error().Err(errServ).Msg("An error ocurred")
	}
}