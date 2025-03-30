package routes

import (
	"api-messages/api/handlers"
	"net/http"
)

func SetupRoutes() {
	handlerInit := handlers.HandlersManage{}
	http.HandleFunc("POST /message", handlerInit.CreateMessage)
	http.HandleFunc("GET /message", handlerInit.GetMessage)
}