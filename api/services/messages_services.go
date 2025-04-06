package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

type GetterMessageServices interface {
	GetMessage() ([]byte, error)
}

type WriterMessageServices interface {
	CreateMessage(*http.Request) (string, error)
}

type ServiceMessageImpl struct{}

func NewServicesImpl() *ServiceMessageImpl {
	return &ServiceMessageImpl{}
}

func (g *ServiceMessageImpl) GetMessage() ([]byte, error) {
	responseMessage, err := json.Marshal("Sucess")
	if err != nil {
		return nil, errors.New("Failed to enconded json response")
	}

	return responseMessage, nil
}

func (w *ServiceMessageImpl) CreateMessage(r *http.Request) (string, error) {
	var message Message
	dataBody, err := io.ReadAll(r.Body)
	if err != nil {
		return "", errors.New("Failed to read body create message")
	}

	errJson := json.Unmarshal(dataBody, &message)
	if errJson != nil {
		return "", errors.New("Failed to parsed json to struct data message")
	}

	fmt.Println(message)
	return message.Message, nil
}
