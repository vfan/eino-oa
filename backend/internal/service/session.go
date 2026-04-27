package service

import (
	"github.com/google/uuid"
	"github.com/wuhaitao/eino-oa/backend/internal/model"
	"github.com/wuhaitao/eino-oa/backend/internal/repository"
)

type SessionService struct {
	repo *repository.SessionRepo
}

func NewSessionService(repo *repository.SessionRepo) *SessionService {
	return &SessionService{repo: repo}
}

func (s *SessionService) List() ([]model.ChatSession, error) {
	return s.repo.List()
}

func (s *SessionService) GetByID(id string) (*model.ChatSession, error) {
	return s.repo.GetByID(id)
}

func (s *SessionService) Create(title string) (*model.ChatSession, error) {
	session := &model.ChatSession{
		ID:    uuid.New().String(),
		Title: title,
	}
	if err := s.repo.Create(session); err != nil {
		return nil, err
	}
	return session, nil
}

func (s *SessionService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *SessionService) AddMessage(sessionID, role, content string) (*model.ChatMessage, error) {
	msg := &model.ChatMessage{
		SessionID: sessionID,
		Role:      role,
		Content:   content,
	}
	if err := s.repo.AddMessage(msg); err != nil {
		return nil, err
	}
	_ = s.repo.Touch(sessionID)
	return msg, nil
}

func (s *SessionService) GetMessages(sessionID string) ([]model.ChatMessage, error) {
	return s.repo.GetMessages(sessionID)
}

func (s *SessionService) UpdateTitle(id, title string) error {
	return s.repo.UpdateTitle(id, title)
}
