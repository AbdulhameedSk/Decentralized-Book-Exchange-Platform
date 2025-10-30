package main

import (
    "decentralized-book-exchange/internal/ledger"
    "decentralized-book-exchange/internal/storage"

    "github.com/gin-gonic/gin"
)

func main() {
    dsn := "host=localhost user=postgres password=yourpassword dbname=bookexchange port=5432 sslmode=disable"
    store := storage.NewDBStorage(dsn)
    router := gin.Default()

    // Add a new user
    router.POST("/users", func(c *gin.Context) {
        var user ledger.User
        if err := c.BindJSON(&user); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        if err := store.AddUser(user); err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, user)
    })

    // List all users
    router.GET("/users", func(c *gin.Context) {
        users, err := store.ListUsers()
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, users)
    })

    // Add & Get Books Endpoints (similar)
    router.POST("/books", func(c *gin.Context) {
        var book ledger.Book
        if err := c.BindJSON(&book); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        // verify owner exists
        if _, err := store.GetUserByID(book.OwnerID); err != nil {
            c.JSON(404, gin.H{"error": "USER ID INCORRECT"})
            return
        }
        book.Available = true
        book.AddedBy = book.OwnerID
        if err := store.AddBook(book); err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, book)
    })

    router.GET("/books", func(c *gin.Context) {
        books, err := store.ListBooks()
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, books)
    })

    //Get Book by ID
    router.GET("/books/:id", func(c *gin.Context) {
        id := c.Params.ByName("id")
        book, err := store.GetBookByID(id)
        if err != nil {
            c.JSON(404, gin.H{"error": "Book not found"})
            return
        }
        c.JSON(200, book)
    })

    //Get User by ID
    router.GET("/users/:id", func(c *gin.Context) {
        id := c.Params.ByName("id")
        user, err := store.GetUserByID(id)
        if err != nil {
            c.JSON(404, gin.H{"error": "USER DOES NOT EXISTS"})
            return
        }
        c.JSON(200, user)
    })

    // Start the server
    router.Run(":8080")
}
