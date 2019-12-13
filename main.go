package main

import (
	"github.com/Muha113/homework3/server"
)

func main() {
	srv := server.NewHTTPServer("", "8083")

	srv.Serve()
}
