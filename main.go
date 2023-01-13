package main

import (
	" hery-ciaputra/assignment-05-golang-backend/db"
	" hery-ciaputra/assignment-05-golang-backend/server"
	"log"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Println("Failed to connect to database: ", err)
	}

	server.Init()
}
