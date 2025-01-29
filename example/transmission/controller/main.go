package main

import (
	"bufio"
	"fmt"
	"github.com/hasuburero/echonetlite/echonetlite/controller"
	"os"
)

const (
	multicastaddr = "224.0.23.0"
	multicastport = 3610
	unicastaddr   = "localhost"
	unicastport   = 3611
)

func main() {
	cont_instance, err := controller.Start(multicastaddr, multicastport, unicastaddr, unicastport)
	if err != nil {
		fmt.Println(err)
		fmt.Println("controller.Start error")
		return
	}

	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("<< ")
		scan.Scan()
		fmt.Println(scan.Text())
		err = cont_instance.Send(scan.Bytes())
		if err != nil {
			fmt.Println(err)
			fmt.Println("Controller_instance.Send error")
			continue
		}
		frame, err := cont_instance.Read()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Controller_instance.Read error")
			continue
		}

		fmt.Println(string(frame))
	}
}
