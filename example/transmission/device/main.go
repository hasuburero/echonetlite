package main

import (
	"fmt"
	"github.com/hasuburero/echonetlite/echonetlite/device"
)

const (
	multicastaddr = "224.0.23.0"
	multicastport = 3610
	unicastport   = 3611
)

func main() {
	dev_instance, err := device.Start(multicastaddr, multicastport, unicastport)
	if err != nil {
		fmt.Println(err)
		fmt.Println("device.Start error")
		return
	}

	for {
		srcIP, frame, err := dev_instance.Read()
		if err != nil {
			fmt.Println(err)
			fmt.Println("device.Device_instance.Read error")
			continue
		}
		fmt.Println(string(frame))
		err = dev_instance.Send([]byte("Ack"), srcIP)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Device_instance.Send error")
			continue
		}
	}
}
