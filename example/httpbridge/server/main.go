package main

import (
	"fmt"
	"github.com/hasuburero/echonetlite/echonetlite"
	"github.com/hasuburero/echonetlite/echonetlite/class/bat"
	"github.com/hasuburero/echonetlite/httpbridge/server"
	"http"
	"os"
	"strconv"
)

const (
	GW_num = 1
)

var (
	Class_VGW = [3]byte{0x10, 0xff, 0x01}
	EHD1      = echonetlite.EHD1
	EHD2      = echonetlite.EHD2
)

const (
	addr               = ""
	port               = "8080"
	contract_path      = "/contract"
	data_path          = "/data"
	StatusOK           = 200
	StatusUnauthorized = 401
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
	GW = make(map[string]GW_struct)
	for i := range GW_num {
		gw_id := strconv.Itoa(i)
		GW[gw_id] = GW_struct{Gw_id: gw_id, Tid: [2]byte{0, 0}}
	}

	Bridge_instance := server.Init(addr, port, contract_path, data_path)
	fmt.Println("echonetlite bridge server started")
	for {
		select {
		case contract_data := <-Bridge_instance.Read_recv_contract:
			fmt.Println("receive contract")
			// here is good place to verify gw_id
			gw_id := contract_data.Get_contract_request.Gw_id
			_, exists := GW[gw_id]
			if !exists {
				contract_data.Return_channel <- server.ReturnChannel{StatusCode: http.StatusNotAcceptable}
				continue
			}
			var instance echonetlite.Echonetlite
			instance = echonetlite.Echonetlite{EHD1: EHD1, EHD2: EHD2, Tid: GW[gw_id].Tid,
				SEOJ: Class_VGW, DEOJ: echonetlite.Class_Battery, ESV: echonetlite.ESV_Get,
				OPC: byte(1)}
			instance.Datactx = append(instance.Datactx, echonetlite.Datactx{EPC: bat.SOC.EPC, PDC: bat.SOC.Size})
			err := instance.MakeFrame()
			if err != nil {
				fmt.Println(err)
				fmt.Println("echonetlite.Echonetlite.MakeFrame error")
				os.Exit(1)
			}
			for _, ctx := range []byte(instance.Frame) {
				fmt.Printf("%x ", ctx)
			}

			fmt.Println("")
			instance.ShowInstanceFrame()
			contract_data.Return_channel <- server.ReturnChannel{Echonet_instance: instance, StatusCode: StatusOK}
		case data_data := <-Bridge_instance.Read_recv_data:
			fmt.Println("receive data")
			//gw_id := data_data.Post_data_request.Gw_id
			frame := data_data.Post_data_request.Frame

			echonet_instance := echonetlite.MakeInstance([]byte(frame))
			echonetlite.ShowByteFrame([]byte(frame))
			fmt.Print("data:")
			echonet_instance.ShowInstanceFrame()
		}
	}
	return
}
