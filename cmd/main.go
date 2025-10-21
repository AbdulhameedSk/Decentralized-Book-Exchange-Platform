package main

import (
    "github.com/gin-gonic/gin"
    "decentralized-book-exchange/internal/storage"
    "decentralized-book-exchange/internal/ledger"
)

func main() {
    store := storage.NewStorage()
    router := gin.Default()

    // Add a new user
    router.POST("/users", func(c *gin.Context) {
        var user ledger.User
        if err := c.BindJSON(&user); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        store.AddUser(user)
        c.JSON(200, user)
    })

    // List all users
    router.GET("/users", func(c *gin.Context) {
        c.JSON(200, store.ListUsers())
    })

    // Add & Get Books Endpoints (similar)
    router.POST("/books", func(c *gin.Context) {
        var book ledger.Book
        if err := c.BindJSON(&book); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        store.AddBook(book)
        c.JSON(200, book)
    })

    router.GET("/books", func(c *gin.Context) {
        c.JSON(200, store.ListBooks())
    })

    router.Run(":8080")
}
