/*
Battery class definition
*/
package bat

import (
	"github.com/hasuburero/echonetlite/echonetlite/class"
)

const (
	ClassGroupCode = 0x02 // 1byte, クラスグループコード
	ClassCode      = 0x7d // 1byte, クラスコード
	EPC_Status     = 0x80 // 1byte, 動作状態 on / off
	EPC_Identifier = 0x83 // 9 or 17byte, オブジェクトを固有に識別する番号
	EPC_Exception  = 0x89 // 2byte, 異常内容
	EPC_Product    = 0x8c // 12byte, product code
)

var Status = class.Class{
	Name:       "動作状態",
	EPC:        0x80,
	Context:    "on/offの状態, 0x30/0x31",
	DataType:   "unsigned char",
	Size:       1,
	AccessRule: "Set/Get"}
var Identifier = class.Class{
	Name:       "識別番号",
	EPC:        0x83,
	Context:    "オブジェクトを固有に識別する番号",
	DataType:   "unsigned char",
	Size:       9,
	AccessRule: "Get"}
var Error = class.Class{
	Name:       "故障状態",
	EPC:        0x88,
	Context:    "異常あり:0x41, 無し:0x42",
	DataType:   "unsigned char",
	Size:       1,
	AccessRule: "Get"}
var ErrorMSG = class.Class{
	Name:       "エラーメッセージ",
	EPC:        0x89,
	Context:    "上位1byte:異常内容小分類, 下位1byte:大分類, 異常内容プロパティ参照",
	DataType:   "unsigned short",
	Size:       2,
	AccessRule: "Get"}
var Maker = class.Class{
	Name:       "メーカコード",
	EPC:        0x8a,
	Context:    "メーカコード",
	DataType:   "unsigned char",
	Size:       3,
	AccessRule: "Get"}
var Product = class.Class{
	Name:     "商品コード",
	EPC:      0x8c,
	Context:  "ascii, makers definition",
	DataType: "unsigned char", Size: 12, AccessRule: "Get"}
var Time = class.Class{
	Name:       "時刻設定",
	EPC:        0x97,
	Context:    "時刻設定HH:MM, 0-23, 0-59",
	DataType:   "unsigned char",
	Size:       2,
	AccessRule: "Set/Get"}
var Date = class.Class{
	Name:       "現在年月日設定",
	EPC:        0x98,
	Context:    "現在年月日YYYY:MM:DD",
	DataType:   "unsigned char",
	Size:       4,
	AccessRule: "Set/Get"}
var ChargeableCap = class.Class{
	Name:       "充電可能量",
	EPC:        0xa4,
	Context:    "現時点での充電可能量, 0-999,999,999Wh",
	DataType:   "unsigned long",
	Size:       4,
	AccessRule: "Get"}
var DischargeableCap = class.Class{
	Name:       "放電可能量",
	EPC:        0xa5,
	Context:    "現時点での放電可能量, 0-999,999,999Wh",
	DataType:   "unsigned long",
	Size:       4,
	AccessRule: "Get"}
var Forward = class.Class{
	Name:       "積算充電量",
	EPC:        0xa8,
	Context:    "0.001kWh, 0-999,999.999kWh",
	DataType:   "unsigned long",
	Size:       4,
	AccessRule: "Get"}
var Backward = class.Class{
	Name:       "積算放電量",
	EPC:        0xa9,
	Context:    "0.001kWh, 0-999,999.999kWh",
	DataType:   "unsigned long",
	Size:       4,
	AccessRule: "Get"}
var MaxCharge = class.Class{
	Name:       "最大充電電力",
	EPC:        0xc8,
	Context:    "実際は最小:最大の2つを同時取得, 0-999,999,999W",
	DataType:   "unsigned long * 2",
	Size:       8,
	AccessRule: "Get"}
var MaxDischarge = class.Class{
	Name:       "最大放電電力",
	EPC:        0xc9,
	Context:    "実際は最小:最大の2つを同時取得, 0-999,999,999W",
	DataType:   "unsigned long * 2",
	Size:       8,
	AccessRule: "Get"}
var Mode_current = class.Class{
	Name:       "運転動作状態",
	EPC:        0xcf,
	Context:    "急速充電:0x41, 充電:0x42, 放電:0x43, 待機:0x44, テスト:0x45, 自動:0x46, 再起動:0x48, 実効容量再計算処理:0x49, その他:0x40",
	DataType:   "unsigned char",
	Size:       1,
	AccessRule: "Get"}
var Size = class.Class{
	Name:       "DC定格電力量",
	EPC:        0xd0,
	Context:    "0-999,999,999Wh",
	DataType:   "unsigned long",
	Size:       4,
	AccessRule: "Get"}
var GenPower = class.Class{
	Name:       "瞬時電力",
	EPC:        0xd3,
	Context:    "+-, 1-999,999,999W",
	DataType:   "signed long",
	Size:       4,
	AccessRule: "Get"}
var Mode_set = class.Class{
	Name:       "運転モード設定値",
	EPC:        0xda,
	Context:    "急速充電:0x41, 充電:0x42, 放電:0x43, 待機:0x44, テスト:0x45, 自動:0x46, 再起動:0x48, 実効容量再計算処理:0x49, その他:0x40",
	DataType:   "unsigned char",
	Size:       1,
	AccessRule: "Set/Get"}
var SOC = class.Class{
	Name:       "蓄電池SOC",
	EPC:        0xe2,
	Context:    "蓄電池残量, 0-999,999,999Wh",
	DataType:   "unsigned long",
	Size:       4,
	AccessRule: "Get"}
