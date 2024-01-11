package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"software_containerization/Models"
)

func main() {
	// set the flags for the logging package to give the filename in the logs
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("starting sever")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(w, "Hello World")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
	database_int()

}

// func database_int() gorm.DB { // not sure if needed
func database_int() {
	db_url := os.Getenv("DATABASE_URL")
	if db_url == "" {
		log.Fatal("DATABASE_URL environment variable not set")
		os.Exit(-1)
	}
	dsn := "host=database user=postgres password=postgres dbname=Events port=9920 sslmode=disable TimeZone=Europe/Berlin"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connected")
	}

	err = db.AutoMigrate(&Models.Event{}, &Models.User{}, &Models.Role{})

	if err == nil {
		log.Fatal(err)
	}

	//return gorm.DB{}
	//return *db // not sure pointer arithmetics
}
