package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"software_containerization/Models"
	"time"
)

func main() {
	// set the flags for the logging package to give the filename in the logs
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("starting sever")
	database_int()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(w, "Hello World")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

}

// func database_int() gorm.DB { // not sure if needed
func database_int() {
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
	//return *db // not sure pointer arithmetics
}
