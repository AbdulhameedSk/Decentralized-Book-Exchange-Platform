package ledger

type User struct {
    ID       string `gorm:"primaryKey"`
    Name     string `gorm:"not null"`
    Email    string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    Books    []Book `gorm:"foreignKey:OwnerID"`
}

type Book struct {
    ID        string `gorm:"primaryKey"`
    Title     string `gorm:"not null"`
    Author    string
    OwnerID   string
    Available bool
    AddedBy   string
}
