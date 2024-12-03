package main

import (
	"fmt"
	"github.com/hasuburero/echonetlite/echonetlite"
	"github.com/hasuburero/echonetlite/httpbridge/server"
	"os"
	"strconv"
)

const (
	GW_num = 2
)

var (
	Class_VGW = [3]byte{0x10, 0xff, 0x01}
	EHD1      = echonetlite.EHD1
	EHD2      = echonetlite.EHD2
)

const (
	addr          = ""
	port          = ":8080"
	contract_path = "/contract"
	data_path     = "/data"
)

const (
	SetI = echonetlite.ESV_SetI
	SetC = echonetlite.ESV_SetC
	Get  = echonetlite.ESV_Get
)

var frame []byte = []byte{}

type GW_struct struct {
	Gw_id string
	Tid   [2]byte
}

var GW map[string]GW_struct

func main() {
	for i := range GW_num {
		gw_id := strconv.Itoa(i)
		GW[gw_id] = GW_struct{Gw_id: gw_id, Tid: [2]byte{0, 0}}
	}

	Bridge_instance := server.Init(addr, port, contract_path, data_path)
	for {
		select {
		case contract_data := <-Bridge_instance.Read_recv_contract:
			gw_id := contract_data.Get_contract_request.Gw_id
			fmt.Println(gw_id)
			var instance echonetlite.Echonetlite
			instance = echonetlite.Echonetlite{EHD1: EHD1, EHD2: EHD2, Tid: GW[gw_id].Tid,
				SEOJ: Class_VGW, DEOJ: echonetlite.Class_SmartMeter, ESV: echonetlite.ESV_Get,
				OPC: byte(0), Datactx: []echonetlite.Datactx{}}
			err := instance.MakeFrame()
			if err != nil {
				fmt.Println(err)
				fmt.Println("echonetlite.Echonetlite.MakeFrame error")
				os.Exit(1)
			}
			contract_data.Return_channel <- instance
		case data_data := <-Bridge_instance.Read_recv_data:
			gw_id := data_data.Post_data_request.Gw_id
			frame := data_data.Post_data_request.Frame
			fmt.Println(gw_id)
			fmt.Println(frame)
			var echonetlite = echonetlite.Echonetlite{Frame: []byte(frame)}
			err := echonetlite.ReverseFrame()
			if err != nil {
				fmt.Println(err)
				fmt.Println("echonetlite.Echonetlite.ReverseFrame error")
				os.Exit(1)
			}
			fmt.Println(echonetlite)
		}
	}

	return
}
