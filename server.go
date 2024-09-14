package main

import (
    "github.com/gin-gonic/gin"
    "github.com/glebarez/sqlite"
    "gorm.io/gorm"
    "fmt"
)

type Brainee struct {
    gorm.Model
    Text   string `json:"text"`
    Author string `json:"author"`
    Brand  string `json:"brand"`
}

func main() {
    router := gin.Default()

    db, err := gorm.Open(sqlite.Open("brainees.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&Brainee{})

    router.POST("/brainees", func(c *gin.Context) {
        var brainee Brainee

        if err := c.ShouldBindJSON(&brainee); err != nil {
            c.JSON(400, gin.H{"error": "Invalid JSON data"})
            return
        }

        db.Create(&brainee)

        c.JSON(200, brainee)
    })

    router.GET("/brainees/:braineeId", func(c *gin.Context) {
        var brainee Brainee

        id := c.Param("id")
        c.String(200, "Getting a Brainee\n")

        result := db.First(&brainee, id)
        if result.Error != nil {
            c.JSON(404, gin.H{"error": "Brainee not found"})
            return
        }

        fmt.Printf("Retrieved Brainee: ID: %s, Text: %s, Author: %s, Brand: %s\n", 
            id, brainee.Text, brainee.Author, brainee.Brand)

        c.JSON(200, brainee)
    })

    router.Run(":8080")
}
