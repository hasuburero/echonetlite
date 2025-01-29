package main

import (
	"fmt"
	"github.com/hasuburero/echonetlite/echonetlite"
	"github.com/hasuburero/echonetlite/echonetlite/class"
	"github.com/hasuburero/echonetlite/echonetlite/class/bat"
)

var EHD1 = byte(0x10)
var EHD2 = byte(0x10)
var Tid = [2]byte{0x00, 0x00}
var Seoj = [3]byte{0x0e, 0xf0, 0x00}
var Deoj = [3]byte{0x00, 0x00, 0x00}
var ESV_Get = byte(0x62)
var ESV_Get_Res = byte(0x72)
var OPC = byte(14)

func makeechonetlite() []byte {
	var epc_map = make(map[byte]class.Class)
	epc_map[0x88] = bat.Error
	epc_map[0x89] = bat.ErrorMSG
	epc_map[0x8a] = bat.Maker
	epc_map[0xa4] = bat.ChargeableCap
	epc_map[0xa5] = bat.DischargeableCap
	epc_map[0xa8] = bat.Forward
	epc_map[0xa9] = bat.Backward
	epc_map[0xc8] = bat.MaxCharge
	epc_map[0xc9] = bat.MaxDischarge
	epc_map[0xcf] = bat.Mode_current
	epc_map[0xd0] = bat.Size
	epc_map[0xd3] = bat.GenPower
	epc_map[0xda] = bat.Mode_set
	epc_map[0xe2] = bat.SOC

	var byte_buf = []byte{}
	byte_buf = append(byte_buf, EHD1)
	byte_buf = append(byte_buf, EHD2)
	byte_buf = append(byte_buf, Tid[:]...)
	byte_buf = append(byte_buf, Seoj[:]...)
	byte_buf = append(byte_buf, Deoj[:]...)
	byte_buf = append(byte_buf, ESV_Get)
	byte_buf = append(byte_buf, OPC)
	for _, ctx := range epc_map {
		byte_buf = append(byte_buf, ctx.EPC)
		byte_buf = append(byte_buf, ctx.Size)
	}
	return byte_buf
}

func main() {
	testframe := makeechonetlite()
	instance, err := echonetlite.MakeInstance(testframe)
	if err != nil {
		fmt.Println(err)
		fmt.Println("echonetlite.Makeinstance error")
		return
	}
	fmt.Println(instance)
}
