package repository

import (
	"database/sql"
)

type MessagesData struct {
	IdMsg int64 `db:"id" json:"id_message"`
	Content string `db:"content" json:"message"`
}

type WriterMessageRepository interface {
	CreateMessage(content string) (bool, error)
	UpdateMessage(content string, id int64) (bool, error)
	DeleteMessage(id int64) (bool, error)
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
	rows, err := r.db.Query("SELECT s.id, s.content FROM messages as s ORDER BY s.id ASC")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var message MessagesData
		err = rows.Scan(&message.IdMsg, &message.Content)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
} 

func (r *MessageRepoImpl) UpdateMessage(content string, id int64) (bool, error) {
	_, err := r.db.Exec("UPDATE messages as s SET content = $1 WHERE id = $2", content, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *MessageRepoImpl) DeleteMessage(id int64) (bool, error) {
	_, err := r.db.Exec("DELETE FROM messages as ms where ms.id = $1", id)
	if err != nil {
		return false, err
	}

	return true, nil
}