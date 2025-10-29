package ledger

type Book struct {
	ID        string
	Title     string
	Author    string
	OwnerID   string
	Available bool
	AddedBy   string
}

type User struct {
	ID       string
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
	// To export capital B in Books
	Books []Book
}
