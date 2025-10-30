package ledger

type Book struct {
	ID        string `gorm:"primaryKey"`
	Title     string
	Author    string
	OwnerID   string
	Available bool
	AddedBy   string
}

type User struct {
	ID       string `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
	// To export capital B in Books
	Books []Book `gorm:"foreignKey:OwnerID"`
}
