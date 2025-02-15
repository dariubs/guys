package main

import (
	"log"

	"github.com/dariubs/guys/app/ctrl/index"
	"github.com/dariubs/guys/app/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func main() {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB, err = database.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	index := index.NewIndex(DB)

	router.LoadHTMLGlob("view/index/**/*")

	router.GET("/", index.Index())
	router.GET("/:username", index.Username())

	router.Run(":8080")
}
