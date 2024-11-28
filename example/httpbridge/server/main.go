package main

import (
	"fmt"
	"github.com/hasuburero/echonetlite/httpbridge/server"
)

func main() {
	wait := make(chan bool)
	server.Init()
	<-wait
}
