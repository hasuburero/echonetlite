package main

import (
	"fmt"
	"github.com/hasuburero/echonetlite/echonetlite"
	"github.com/hasuburero/echonetlite/httpbridge/server"
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
const (
	esv = byte()
)

var frame []byte = []byte{}

type GW_struct struct {
	Gw_id string
	Tid   int16
}

var GW map[string]GW_struct

func Data_controller() {
}

func main() {
	for i := range GW_num {
		gw_id := strconv.Itoa(i)
		GW[gw_id] = GW_struct{Gw_id: gw_id, Tid: 0}
	}

	Bridge_instance := server.Init(addr, port, contract_path, data_path)
	for {
		select {
		case contract_data := <-Bridge_instance.Read_recv_contract:
			gw_id := contract_data.Get_contract_request.Gw_id
			fmt.Println(gw_id)
			echonetlite.MakeInstance(EHD1, EHD2, GW[gw_id].Tid++, Class_VGW, echonetlite.Class_SmartMeter)
			contract_data.Return_channel
		case data_data := <-Bridge_instance.Read_recv_data:
		}
	}
}
