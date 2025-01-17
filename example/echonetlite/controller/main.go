package main

import (
	"fmt"
	"github.com/hasuburero/echonetlite/echonetlite/controller"
)

const (
	multicastaddr = "224.0.23.0"
	multicastport = 3610
	unicastaddr = "localhost"
	unicastport = 3610
)
func controller.Start(multicastAddr string, unicastAddr string, multicastPort int, unicastPort int) (controller.Controller_Instance, error)

func main() {
	controller_instance := controller.Start()
}
