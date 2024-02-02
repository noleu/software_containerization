package main

import (
	"fmt"
	"log"
	"os"
	"software_containerization/Handlers"
	"software_containerization/Models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// set the flags for the logging package to give the filename in the logs
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("starting server")
	db := database_init()

	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/api")

	api.POST("/events", Handlers.CreateEventHandler(db))
	api.GET("/events", Handlers.GetEventsHandler(db)) // get all events, get event by id, or by userId query params
	api.DELETE("/events/:id", Handlers.DeleteEventHandler(db))
	api.PUT("/events/:id", Handlers.UpdateEventHandler(db))

	r.Run(":80")
}

func database_init() *gorm.DB {
	log.Println("connecting to database")
	log.Println("waiting for database to start")
	time.Sleep(time.Duration(5) * time.Second)

	dbHostName := os.Getenv("DB_HOST_NAME")
	if dbHostName == "" {
		log.Fatal("Environment variable DB_HOST_NAME is not set")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal("Environment variable DB_USER is not set")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("Environment variable DB_PASSWORD is not set")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("Environment variable DB_NAME is not set")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Berlin",
		dbHostName, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("connection to database established")
	log.Println("migrating database")
	err = db.AutoMigrate(&Models.Event{})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("database migration finished")
	//return gorm.DB{}
	return db
}
