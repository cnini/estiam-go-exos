package main

import (
	"estiam-go-exos/server"
)

func main() {
	server.Start()
	server.ListenAndServe()
}
