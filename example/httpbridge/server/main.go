package main

import (
	"fmt"
	"httpbridge/server"
)

func main() {
	wait := make(chan bool)
	server.Init()
}
