package storage

import (
	"decentralized-book-exchange/internal/ledger"
	"sync"
)
//Simple in memory storage
type Storage struct {
	users map[string]ledger.User
	books map[string]ledger.Book
	// RWMutex stands for Read-Write Mutex helps to manage concurrent access to shared resources.
	mu sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		users: make(map[string]ledger.User),
		books: make(map[string]ledger.Book),
	}
}

func (s *Storage) AddUser(user ledger.User) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.users[user.ID] = user
}

func (s *Storage) ListUsers() []ledger.User {	
	s.mu.RLock()
	defer s.mu.RUnlock()
	users := []ledger.User{}
	for _, u := range s.users {
		users = append(users, u)
	}
	return users
}

func (s *Storage) GetUser(id string) (ledger.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	u, ok := s.users[id]
	return u, ok
}

func (s *Storage) AddBook(book ledger.Book) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.books[book.ID] = book
}

func (s *Storage) ListBooks() []ledger.Book {
	s.mu.RLock()
	defer s.mu.RUnlock()
	books := []ledger.Book{}
	for _, b := range s.books {
		books = append(books, b)
	}
	return books
}

func (s *Storage) GetBook(id string) (ledger.Book, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	b, ok := s.books[id]
	return b, ok
}
