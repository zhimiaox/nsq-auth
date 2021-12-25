package main

import (
	"sync"
)

var (
	storageIns  Storage
	storageOnce sync.Once
)

type Storage interface {
	Refresh()
	Get(secret string) []Authorization
	Set(secret string, auth []Authorization)
}

type storage struct {
	sync.RWMutex
	data map[string][]Authorization
}

func (s *storage) Refresh() {
	s.Lock()
	defer s.Unlock()
	s.data = make(map[string][]Authorization)
	for _, v := range GetPlugins() {
		for secret, auth := range v.Authorization() {
			s.data[secret] = auth
		}
	}
}

func (s *storage) Get(secret string) []Authorization {
	s.RLock()
	defer s.RUnlock()
	return s.data[secret]
}

func (s *storage) Set(secret string, auth []Authorization) {
	s.Lock()
	defer s.Unlock()
	s.data[secret] = auth
}

func GetStorage() Storage {
	storageOnce.Do(func() {
		storageIns = &storage{
			data: make(map[string][]Authorization),
		}
	})
	return storageIns
}
