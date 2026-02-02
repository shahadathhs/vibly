package store

import (
	"fmt"
	"vibly/app/models"
)

type ChatStore struct {
	BaseDir string
}

func (s *ChatStore) getStore(channelID string) *FileStore[models.Message] {
	return &FileStore[models.Message]{
		FilePath: fmt.Sprintf("%s/chat_%s.json", s.BaseDir, channelID),
	}
}

func (s *ChatStore) SaveMessage(channelID string, msg models.Message) error {
	fs := s.getStore(channelID)
	if err := fs.Init(); err != nil {
		return err
	}

	messages, err := fs.Load()
	if err != nil {
		return err
	}

	messages = append(messages, msg)
	return fs.Save(messages)
}

func (s *ChatStore) GetMessages(channelID string, limit int) ([]models.Message, error) {
	fs := s.getStore(channelID)
	// We don't want to error if the file doesn't exist yet, just return empty
	if _, err := fs.Load(); err != nil {
		return []models.Message{}, nil
	}
	
	messages, err := fs.Load()
	if err != nil {
		return nil, err
	}

	if limit > 0 && len(messages) > limit {
		messages = messages[len(messages)-limit:]
	}

	return messages, nil
}
