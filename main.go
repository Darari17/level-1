package main

import (
	"level-1/config"
	"level-1/controller"
	"level-1/database"
	"level-1/repository"
	"level-1/router"
	"level-1/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}
	defer sqlDB.Close()

	err = database.AutoMigrate(db)
	if err != nil {
		log.Fatal("Database migration failed:", err)
	}

	app := gin.Default()
	group := app.Group("/api/v1")

	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)
	bookRoutes := router.NewBookRoutes(bookController)
	bookRoutes.SetupRoutes(group)

	port := ":3000"
	log.Println("Server is running on port", port)
	if err := app.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
