package repository

import (
	"database/sql"
)

type WriterMessageRepository interface {
	CreateMessage(content string) (bool, error)
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