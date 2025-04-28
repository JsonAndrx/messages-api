package services

import (
	"api-messages/api/repository"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

type MessageWithId struct {
	Id int64 `json:"Id"`
	Message string `json:"message"`
}

type GetterMessageServices interface {
	GetMessage() ([]byte, error)
}

type WriterMessageServices interface {
	CreateMessage(*http.Request) (string, error)
	UpdateMessage(*http.Request) (string, error)
	DeleteMessage(*http.Request) (string, error)
}

type repositoryMessage interface {
	repository.WriterMessageRepository
	repository.GetterMessageRepository
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
	dataMsg, err := g.repo.GetMessages()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("failed to get messages")
	}
	
	responseMessage, err := json.Marshal(dataMsg)
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

func (w *ServiceMessageImpl) UpdateMessage(r *http.Request) (string, error) {
	var message MessageWithId
	dataBody, err := io.ReadAll(r.Body)
	if err != nil {
		return "", errors.New("failed to read body create message")
	}

	err = json.Unmarshal(dataBody, &message)
	if err != nil {
		return "", errors.New("failed to parsed json to struct data message")
	}

	updateMessage, err := w.repo.UpdateMessage(message.Message, message.Id)
	if  !updateMessage {
		return "", err
	}

	return "Update message success", nil
}

func (w *ServiceMessageImpl) DeleteMessage(r *http.Request) (string, error) {
	var msg MessageWithId
	dataBody, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(dataBody, &msg)
	if err != nil {
		return "", errors.New("failed to parsen json to struct message")
	}

	deleteMsg, err := w.repo.DeleteMessage(msg.Id)
	if !deleteMsg {
		return "", err
	}

	return "Delete message success", nil
}
