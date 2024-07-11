package main

import (
	"bitespeed/controllers"
	"bitespeed/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=bitespeed sslmode=disable password=Ruchir@123")
    if err != nil {
        panic("failed to connect to database")
    }
    defer db.Close()

    db.AutoMigrate(&models.Contact{})

    r := gin.Default()

    contactController := controllers.NewContactController(db)

    r.POST("/identify", contactController.Identify)

    r.Run()
}
