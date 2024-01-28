package main

import (
	"log"
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

	r.Run(":8080")
}

func database_init() *gorm.DB {
	log.Println("connecting to database")
	log.Println("waiting for database to start")
	time.Sleep(time.Duration(5) * time.Second)

	dsn := "host=database user=postgres password=postgres dbname=Events port=5432 sslmode=disable TimeZone=Europe/Berlin"
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
