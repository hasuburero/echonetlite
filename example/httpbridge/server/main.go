package main

import (
	"fmt"
	"github.com/hasuburero/echonetlite/httpbridge/server"
)

func main() {
	wait := make(chan bool)
	Bridge_instance := server.Init()
	Bridge_instance.
	<-wait
}
