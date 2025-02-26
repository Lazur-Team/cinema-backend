package main

import (
	"fmt"
	"log"
	"server/internal/config"
	"server/internal/db"
	"server/internal/server"
)

func main() {
	config.LoadConfig()
	db.InitDB()

	s := server.NewServer()
	fmt.Println("Server running on port", config.Cfg.ServerPort)
	log.Fatal(s.Start())
}
