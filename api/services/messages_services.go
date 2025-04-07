package services

import (
	"api-messages/api/repository"
	"encoding/json"
	"errors"
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

type repositoryMessage interface {
	repository.WriterMessageRepository
}

type ServiceMessageImpl struct {
	repo repositoryMessage
}

func NewServicesImpl(repo repositoryMessage) *ServiceMessageImpl {
	return &ServiceMessageImpl{
		repo: repo,
	}
}

func (g *ServiceMessageImpl) GetMessage() ([]byte, error) {
	responseMessage, err := json.Marshal("Sucess")
	if err != nil {
		return nil, errors.New("failed to enconded json response")
	}

	return responseMessage, nil
}

func (w *ServiceMessageImpl) CreateMessage(r *http.Request) (string, error) {
	var message Message
	dataBody, err := io.ReadAll(r.Body)
	if err != nil {
		return "", errors.New("failed to read body create message")
	}

	errJson := json.Unmarshal(dataBody, &message)
	if errJson != nil {
		return "", errors.New("failed to parsed json to struct data message")
	}

	createMsg, errMsg := w.repo.CreateMessage(message.Message)
	if !createMsg {
		return "", errMsg
	}

	return "Create message sucess", nil
}
