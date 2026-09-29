package storage

import (
	"mt-bot/internal/domain"
	"sync"
)

type MemoryStorage struct {
	mtx      sync.RWMutex
	sessions map[int64]*domain.UserSession
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		sessions: make(map[int64]*domain.UserSession),
	}
}

func (m *MemoryStorage) Get(chatID int64) (*domain.UserSession, bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	session, ok := m.sessions[chatID]

	return session, ok
}

func (m *MemoryStorage) Set(chatID int64, session *domain.UserSession) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.sessions[chatID] = session
}

func (m *MemoryStorage) Delete(chatID int64) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	delete(m.sessions, chatID)
}
