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

func recvFrame(frame string) (string, error) {
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
		for i, _ := range echonet_instance.Datactx {
			echonet_instance.Datactx[i].EDT = byte_buf[:echonet_instance.Datactx[i].PDC]
		}
		echonet_instance.ESV = echonetlite.ESV_Get_Res
	case echonetlite.ESV_SetC:
		buf = echonet_instance.SEOJ
		echonet_instance.SEOJ = echonet_instance.DEOJ
		echonet_instance.DEOJ = buf
		for i, _ := range echonet_instance.Datactx {
			echonet_instance.Datactx[i].EDT = byte_buf[:echonet_instance.Datactx[i].PDC]
		}
		echonet_instance.ESV = echonetlite.ESV_Set_Res
	}

	err := echonet_instance.MakeFrame()
	return string(echonet_instance.Frame), err
}

func control(frame string) string { // dammy function
	fmt.Print("> ")
	fmt.Println([]byte(frame))

	frame, err := recvFrame(frame)
	if err != nil {
		fmt.Println(err)
		fmt.Println("recvFrame error")
	}

	fmt.Print("< ")
	fmt.Println([]byte(frame))

	return frame
}

func gw_func(arg int) {
	var gw_id string
	gw_id = fmt.Sprintf("%d", arg)
	gw_instance := client.Init(gw_id, Scheme, Addr, Port, Contract_path, Data_path)
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
