package main

import (
	"homeapp/src/server"
)

func init() {
	server.Init()
}

func main() {
	server.Run()
}
