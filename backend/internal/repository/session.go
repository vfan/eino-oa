package repository

import (
	"github.com/wuhaitao/eino-oa/backend/internal/model"
	"gorm.io/gorm"
)

type SessionRepo struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *SessionRepo {
	return &SessionRepo{db: db}
}

func (r *SessionRepo) List() ([]model.ChatSession, error) {
	var sessions []model.ChatSession
	err := r.db.Order("updated_at DESC").Find(&sessions).Error
	return sessions, err
}

func (r *SessionRepo) GetByID(id string) (*model.ChatSession, error) {
	var s model.ChatSession
	err := r.db.Preload("Messages", func(db *gorm.DB) *gorm.DB {
		return db.Order("chat_messages.id ASC")
	}).First(&s, "id = ?", id).Error
	return &s, err
}

func (r *SessionRepo) Create(s *model.ChatSession) error {
	return r.db.Create(s).Error
}

func (r *SessionRepo) UpdateTitle(id, title string) error {
	return r.db.Model(&model.ChatSession{}).Where("id = ?", id).Update("title", title).Error
}

func (r *SessionRepo) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&model.ChatSession{}).Error
}

func (r *SessionRepo) AddMessage(msg *model.ChatMessage) error {
	return r.db.Create(msg).Error
}

func (r *SessionRepo) GetMessages(sessionID string) ([]model.ChatMessage, error) {
	var msgs []model.ChatMessage
	err := r.db.Where("session_id = ?", sessionID).Order("id ASC").Find(&msgs).Error
	return msgs, err
}

func (r *SessionRepo) Touch(id string) error {
	return r.db.Model(&model.ChatSession{}).Where("id = ?", id).UpdateColumn("updated_at", gorm.Expr("CURRENT_TIMESTAMP")).Error
}
