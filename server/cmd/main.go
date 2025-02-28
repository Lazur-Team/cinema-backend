package main

import (
	"log"
	"server/internal/config"
	"server/internal/db"
	"server/internal/server"
)

func main() {
	config.LoadConfig()
	db.InitDB()

	s := server.GetServer()
	log.Fatal(s.Start(server.ServerArgs{DB: db.DB}))
}
