package main

import (
	"server"

	"github.com/painterdrown/go-agenda/entities"
)

func main() {
	entities.InitDB("/data/go-agenda.db")
	server.Start()
}
