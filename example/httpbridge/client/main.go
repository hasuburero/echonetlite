package main

import (
	"fmt"
	"github.com/hasuburero/echonetlite/echonetlite"
	"github.com/hasuburero/echonetlite/httpbridge/client"
	"time"
)

const (
	gw_num = 2
)

const (
	Scheme        = "http://"
	Addr          = "localhost"
	Port          = ":8080"
	Contract_path = "/contract"
	Data_path     = "/data"
)

var wait chan bool

func recvFrame(frame string) string {
	echonet_instance := echonetlite.MakeInstance([]byte(frame))
	var buf [3]byte
	var byte_buf []byte
	for i := range 4 {
		byte_buf = append(byte_buf, byte(i))
	}
	switch echonet_instance.ESV {
	case echonetlite.ESV_Get:
		buf = echonet_instance.SEOJ
		echonet_instance.SEOJ = echonet_instance.DEOJ
		echonet_instance.DEOJ = buf
	case echonetlite.ESV_SetC:
		buf = echonet_instance.SEOJ
		echonet_instance.SEOJ = echonet_instance.DEOJ
		echonet_instance.DEOJ = buf
	case echonetlite.ESV_Get_Res:
		for i, _ := range echonet_instance.Datactx {
			echonet_instance.Datactx[i].EDT = byte_buf[:echonet_instance.Datactx[i].PDC]
		}
	case echonetlite.ESV_Set_Res:
		for i, _ := range echonet_instance.Datactx {
			echonet_instance.Datactx[i].EDT = byte_buf[:echonet_instance.Datactx[i].PDC]
		}
	}

	return
}

func control(frame string) string { // dammy function
	fmt.Print("> ")
	fmt.Println([]byte(frame))

	frame = recvFrame()

	echonet_instance := echonetlite.MakeInstance(frame)

	fmt.Print("< ")
	fmt.Println([]byte(frame))

	return frame
}

func gw_func(arg int) {
	var gw_id string
	fmt.Sprintf(gw_id, "%d", arg)
	gw_instance := client.Init(gw_id, Addr, Port, Contract_path, Data_path)
	for {
		frame, err := gw_instance.Contract()
		if err != nil {
			fmt.Println(err)
			fmt.Println("client.Contract error")
			time.Sleep(time.Second * 2)
			continue
		}
		fmt.Println(string(frame))
		go func() {
			frame = control(frame)
			gw_instance.Data(frame)
		}()
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
