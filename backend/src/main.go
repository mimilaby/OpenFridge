// main package contains the main function
package main

import (
	"openfridge/src/server"
)

func init() {
	server.Init()
}

func main() {
	server.Run()
}
