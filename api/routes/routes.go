package routes

import (
	"api-messages/api/handlers"
	"api-messages/api/repository"
	"api-messages/api/services"
	"database/sql"
	"net/http"
)

func SetupRoutes(db *sql.DB) {
	repoInit := repository.NewRepositoryImpl(db)
	serviceInit := services.NewServicesImpl(repoInit)
	handlerInit := handlers.NewMessageHandler(serviceInit)
	http.HandleFunc("POST /message", handlerInit.CreateMessage)
	http.HandleFunc("GET /message", handlerInit.GetMessage)
	http.HandleFunc("PUT /message", handlerInit.UpdateMessage)
	http.HandleFunc("DELETE /message", handlerInit.DeleteMessage)
}