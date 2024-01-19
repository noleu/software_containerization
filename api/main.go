package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"software_containerization/Models"
	"software_containerization/Handlers"
	"time"
	"github.com/gin-gonic/gin"
)

func InitializeRoles(db *gorm.DB) {
	roles := []string{"Organizer", "Participant", "Guest"}
	for _, roleName := range roles {
		var role Models.Role
		result := db.First(&role, "name = ?", roleName)
		if result.Error == gorm.ErrRecordNotFound {
			// Role not found, create it
			newRole := Models.Role{Name: roleName}
			db.Create(&newRole)
			log.Printf("Role '%s' created.\n", roleName)
		}
	}
}

func main() {
	// set the flags for the logging package to give the filename in the logs
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("starting sever")
	db := database_init()

	r := gin.Default()

	// Initialize roles
	InitializeRoles(db)

	r.GET("/roles", Handlers.GetRolesHandler(db))
	r.GET("/roles/:id", Handlers.GetRoleByIdHandler(db))

	r.POST("/user", Handlers.CreateUserHandler(db))
	r.GET("/users", Handlers.GetUsersHandler(db)) // get all users, get user by id, or by email query params
	r.PUT("/users/:id", Handlers.UpdateUserHandler(db))

	r.POST("/events", Handlers.CreateEventHandler(db))
	r.GET("/events", Handlers.GetEventsHandler(db)) // get all events, get event by id, or by userId query params
	r.DELETE("/events/:id", Handlers.DeleteEventHandler(db))
	r.PUT("/events/:id", Handlers.UpdateEventHandler(db))

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
	err = db.AutoMigrate(&Models.Event{}, &Models.User{}, &Models.Role{})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("database migration finished")
	//return gorm.DB{}
	return db
}
