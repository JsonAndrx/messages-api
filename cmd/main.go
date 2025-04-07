package main

import (
	"api-messages/api/routes"

	"net/http"

	"api-messages/internal"

	"github.com/rs/zerolog/log"
)

func main() {
	connectDb := internal.NewDbImpl()
	db, errDb := connectDb.ConnectDb()
	if errDb != nil {
		log.Error().Err(errDb).Msg("An error ocurred")
	}

	routes.SetupRoutes(db)
	errServ := http.ListenAndServe(":8080", nil)
	if errServ != nil {
		log.Error().Err(errServ).Msg("An error ocurred")
	}
}
