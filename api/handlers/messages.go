package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HandlersManage struct{}

type Message struct {
	Message string `json:"message"`
}

func (h *HandlersManage) GetMessage(w http.ResponseWriter, r *http.Request) {
	responseMessage, err := json.Marshal("Sucess")
	if err != nil {
		http.Error(w, "Failed to enconded json response", http.StatusInternalServerError)
		return
	}
	w.Write(responseMessage)
}

func (h *HandlersManage) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var Message Message
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body to create message", http.StatusInternalServerError)
		return
	}

	errJ := json.Unmarshal(bodyByte, &Message)
	if errJ != nil {
		http.Error(w, "Faile to format response body struct", http.StatusInternalServerError)
		return
	}

	fmt.Println(Message)
	w.Write(bodyByte)
}