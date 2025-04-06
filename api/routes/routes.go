package routes

import (
	"api-messages/api/handlers"
	"api-messages/api/services"
	"net/http"
)

func SetupRoutes() {
	serviceInit := services.NewServicesImpl()
	handlerInit := handlers.NewMessageHandler(serviceInit)
	http.HandleFunc("POST /message", handlerInit.CreateMessage)
	http.HandleFunc("GET /message", handlerInit.GetMessage)
}