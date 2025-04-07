package repository

import (
	"database/sql"
)

type MessagesData struct {
	Content string `db:"content" json:"message"`
}

type WriterMessageRepository interface {
	CreateMessage(content string) (bool, error)
}

type GetterMessageRepository interface {
	GetMessages() ([]MessagesData, error)
}

type MessageRepoImpl struct {
	db *sql.DB
}

func NewRepositoryImpl(db *sql.DB) *MessageRepoImpl {
	return &MessageRepoImpl{
		db: db,
	}
}

func (r *MessageRepoImpl) CreateMessage(content string) (bool, error) {
	_, err := r.db.Exec("INSERT INTO messages (content) VALUES ($1)", content)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *MessageRepoImpl) GetMessages() ([]MessagesData, error) {
	var messages []MessagesData
	rows, err := r.db.Query("SELECT s.content FROM messages as s")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var message MessagesData
		err = rows.Scan(&message.Content)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
} 