package server

import (
	"github.com/urfave/negroni"
)

// Start server
func Start() {
	n := negroni.Classic()
	n.UseHandler(router())
	n.Run("127.0.0.1:8080")
}
