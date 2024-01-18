package main

import (
	"log"
	"software_containerization/Models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	log.Println("starting server")
	db := database_init()

	// Initialize roles
	InitializeRoles(db)

	r := gin.Default()

	r.GET("/roles", GetRolesHandler(db))

	r.POST("/user", CreateUserHandler(db))

	r.GET("/users", GetUserHandler(db))

	r.Run(":8080")
}

func database_init() *gorm.DB {
	log.Println("connecting to database")
	log.Println("waiting for database to start")
	time.Sleep(time.Duration(5) * time.Second)

	dsn := "host=database user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Berlin"
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
