package main

import (
	"fmt"
	"github.com/hasuburero/echonetlite/echonetlite"
	"github.com/hasuburero/echonetlite/httpbridge/client"
	"time"
)

const (
	gw_num = 1
)

const (
	Scheme        = "http://"
	Addr          = "localhost"
	Port          = ":8080"
	Contract_path = "/contract"
	Data_path     = "/data"
)

var wait chan bool

func recvFrame(frame []byte) ([]byte, error) {
	fmt.Println("--- debug ---")
	echonetlite.ShowByteFrame(frame)
	echonet_instance := echonetlite.MakeInstance(frame)
	fmt.Println("--- debug ---")
	echonet_instance.ShowInstanceFrame()
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
	default:
		fmt.Println("default case")
	}

	// debug
	err := echonet_instance.MakeFrame()
	echonet_instance.ShowInstanceFrame()
	return echonet_instance.Frame, err
}

func control(frame []byte) []byte { // dammy function
	fmt.Print("> ")
	fmt.Println(frame)
	fmt.Print("> ")
	echonetlite.ShowByteFrame(frame)

	frame, err := recvFrame(frame)
	if err != nil {
		fmt.Println(err)
		fmt.Println("recvFrame error")
	}

	fmt.Print("< ")
	echonetlite.ShowByteFrame([]byte(frame))

	return frame
}

func gw_func(arg int) {
	var gw_id string
	gw_id = fmt.Sprintf("%d", arg)
	fmt.Println(gw_id)
	gw_instance := client.Init(gw_id, Scheme, Addr, Port, Contract_path, Data_path)
	for {
		frame, err := gw_instance.Contract()
		if err != nil {
			fmt.Println(err)
			fmt.Println("client.Contract error")
			time.Sleep(time.Second * 2)
			continue
		}

		// test case
		frame = control(frame)
		gw_instance.Data(frame)
		// default case
		//go func() {
		//	frame = control(frame)
		//	gw_instance.Data(frame)
		//}()
		time.Sleep(time.Second * 1)
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
