package main

import (
	"github.com/isaqueveras/rest-api-gin-postgres-gorm/database"
	"github.com/isaqueveras/rest-api-gin-postgres-gorm/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
