package main

import "github.com/Muha113/simpleServer/server"

func main() {
	srv := server.NewHTTPServer("", "8083")
	srv.Serve()
}
