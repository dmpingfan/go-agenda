package main

import (
	"github.com/painterdrown/go-agenda/cli/cmd"
	"github.com/painterdrown/go-agenda/entities"
)

func main() {
	entities.InitDB("/data/go-agenda.db")
	cmd.Execute()
}
