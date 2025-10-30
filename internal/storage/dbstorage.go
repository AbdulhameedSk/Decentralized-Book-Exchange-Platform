package storage

import (
	"decentralized-book-exchange/internal/ledger"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Persistent database-backed storage
type DBStorage struct {
	DB *gorm.DB
}

// Initialize PostgreSQL connection and auto-migrate schema
func NewDBStorage(dsn string) *DBStorage {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&ledger.User{}, &ledger.Book{})
	return &DBStorage{DB: db}
}

// Add a new user
func (s *DBStorage) AddUser(user ledger.User) error {
	return s.DB.Create(&user).Error
}

// List all users
func (s *DBStorage) ListUsers() ([]ledger.User, error) {
	var users []ledger.User
	// Preload owned books
	result := s.DB.Preload("Books").Find(&users)
	return users, result.Error
}

// Get user by ID
func (s *DBStorage) GetUserByID(id string) (ledger.User, error) {
	var user ledger.User
	result := s.DB.Preload("Books").First(&user, "id = ?", id)
	return user, result.Error
}

// Add a new book and associate with user
func (s *DBStorage) AddBook(book ledger.Book) error {
	return s.DB.Create(&book).Error
}

// List all books
func (s *DBStorage) ListBooks() ([]ledger.Book, error) {
	var books []ledger.Book
	result := s.DB.Find(&books)
	return books, result.Error
}

// Get book by ID
func (s *DBStorage) GetBookByID(id string) (ledger.Book, error) {
	var book ledger.Book
	result := s.DB.First(&book, "id = ?", id)
	return book, result.Error
}
