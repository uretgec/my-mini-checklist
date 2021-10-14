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

// Database sync from memory to db file if possible
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

// Database import from db file to memory if possible
func (s *Store) Load(db *os.File) {
	json.NewDecoder(db).Decode(&store.Items)
}

// Set new key and value
func (s *Store) Set(key string, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Items[key] = val
}

// Get returns key value
func (s *Store) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.Items[key]
	if !ok {
		return ""
	}

	return val
}

// Del key and value if key exists
func (s *Store) Del(key string) {
	s.mu.Lock()
	if _, ok := s.Items[key]; ok {
		delete(s.Items, key)
	}
	s.mu.Unlock()
}

// GetAll returns all keys and values
func (s *Store) GetAll() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.Items
}

// Stats returns keys count
func (s *Store) Stats() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.Items)
}

// Flush database only inmemory
func (s *Store) Flush() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Items = make(map[string]string)
}
