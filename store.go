package main

import (
	"encoding/json"
	"os"
	"sync"
)

type Store struct {
	mu    sync.RWMutex
	Items map[string]string
}

func NewStore() *Store {
	return &Store{Items: make(map[string]string)}
}

func (s *Store) Sync(db *os.File, firstTime bool) error {
	if firstTime {
		s.mu.RLock()
		defer s.mu.RUnlock()
	}

	err := json.NewEncoder(db).Encode(&s.Items)
	if err != nil {
		return err
	}

	return db.Sync()
}

func (s *Store) Load(db *os.File) {
	json.NewDecoder(db).Decode(&store.Items)
}

func (s *Store) Set(key string, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Items[key] = val
}

func (s *Store) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.Items[key]
	if !ok {
		return ""
	}

	return val
}

func (s *Store) Del(key string) {
	s.mu.Lock()
	if _, ok := s.Items[key]; ok {
		delete(s.Items, key)
	}
	s.mu.Unlock()
}

func (s *Store) GetAll() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.Items
}

func (s *Store) Stats() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.Items)
}

func (s *Store) Flush() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Items = make(map[string]string)
}
