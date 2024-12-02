package main

import (
	"fmt"
	"github.com/hasuburero/echonetlite/httpbridge/client"
	"time"
)

const (
	gw_num = 2
)

const (
	Addr          = ""
	Port          = ":8080"
	Contract_path = "/contract"
	Data_path     = "/data"
)

var wait chan bool

func gw_func(arg int) {
	var gw_id string
	fmt.Sprintf(gw_id, "%d", arg)
	for {
		frame, err := client.Contract(gw_id, Addr, Port, Contract_path, Data_path)
		if err != nil {
			fmt.Println(err)
			fmt.Println("client.Contract error")
			time.Sleep(time.Second * 2)
			continue
		}
		fmt.Println(string(frame))
	}
}

func main() {
	wait = make(chan bool)
	for i := range gw_num {
		go func(i int) {
			gw_func(i)
		}(i)
	}

	<-wait
	return
}
