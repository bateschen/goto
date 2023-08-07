package store

import "sync"

type urlStore struct {
	urls map[string]string
	mu   sync.RWMutex
}

NewURLStore() *urlStore{
    return &urlStore{
        urls: make(map[string]string)
    }
}

func (s *urlStore) get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

func (s *urlStore) set(key, value string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.urls[key]; ok == false {
		return false
	}
	s.urls[key] = value
	return true
}
