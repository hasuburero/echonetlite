package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"echonetlite/echonetlite/frame"
	"io"
	"net/http"
)

const (
	addr = ""
	port = "8080"
)

// data structure witch through the api channel
type Contract_channel struct{
	get_contract_request Get_contract_request
	return_channel chan string
}

type Data_channel struct{
	post_data_request Post_data_request
}

// http request body structure
type Post_data_request struct{
	gw_id string `json:"gw_id"`
	format string `json:"format"`
}

type Get_contract_request struct {
	gw_id  string `json:"gw_id"`
}

// server instance structure
// including contract channel, data channel, 
type Echonet_bridge struct {
	contract_channel chan Contract_channel
	data_channel chan chan 
	echonetlite frame.Echonetlite
}

type Bridge_interface struct {
	gw_id string
	channel chan 
}

func (self *Echonet_instance)ReadRecv(){
}

func (self *Echonet_instance)Contract(w http.ResponseWriter, r *http.Request) {
	length := r.ContentLength
	reqBody := make([]byte, length)
	var ctx Get_contract_request
	err := json.Unmarshal(reqBody, &ctx)
	if err != nil {
		fmt.Println("json.Unmarshal error")
		return
	}
	self.read_recv_channel <- 

	format := Get_echonet_response{format: "1234657890"}
	json_buf, err := json.Marshal(format)
	if err != nil {
		fmt.Println("json.Marshal error")
		return
	}
	body := bytes.NewBuffer(json_buf)
	_, err = io.Copy(w, body)
	if err != nil {
		fmt.Println("io.Copy error")
		return
	}
}

func (self *Echonet_instance)Data(w http.ResponseWriter, r *http.Request){

}

func Init(addr, port, contract, data string) {
	bridge_chan = make(chan chan Bridge_interface)
	server := http.Server{
		Addr: addr + ":" + port,
	}

	http.HandleFunc(contract, Echonet_bridge.Contract)
	http.HandleFunc(data, Echonet_bridge.Data)

	go server.ListenAndServe()
}
