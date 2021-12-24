package main

import (
	"sync"

	"github.com/nsq-auth/models"
)

type Storage interface {
	Refresh()
	Get(secret string) *models.Authorization
}

type storage struct {
	sync.RWMutex
	data map[string]models.Authorization
}

func (s *storage) Refresh() {
	s.Lock()
	s.data = make(map[string]models.Authorization)
	s.Unlock()
}

func (s *storage) Get(secret string) *models.Authorization {
	s.RLock()
	defer s.RUnlock()
	v := s.data[secret]
	return &v
}

func NewStorage() Storage {
	return &storage{
		data: make(map[string]models.Authorization),
	}
}

func RefreshData() {

}
