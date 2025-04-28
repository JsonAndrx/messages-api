package handlers

import (
	"api-messages/api/services"
	"net/http"
)

type messageServices interface {
	services.GetterMessageServices
	services.WriterMessageServices
}

type HandlersManage struct{
	ServiceMessage messageServices
}

func NewMessageHandler(service messageServices) *HandlersManage {
	return &HandlersManage{
		ServiceMessage: service,
	}
}

func (h *HandlersManage) GetMessage(w http.ResponseWriter, r *http.Request) {
	responseMessage, err := h.ServiceMessage.GetMessage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(responseMessage)
}

func (h *HandlersManage) CreateMessage(w http.ResponseWriter, r *http.Request) {
	resMessage, err := h.ServiceMessage.CreateMessage(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(resMessage))
}

func (h *HandlersManage) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	updateMessageRes, err := h.ServiceMessage.UpdateMessage(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(updateMessageRes))
}

func (h *HandlersManage) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	deleteMsg, err := h.ServiceMessage.DeleteMessage(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(deleteMsg))
}