package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hasuburero/echonetlite/echonetlite"
	"io"
	"net/http"
)

// data structure witch through the api channel
type Contract_context struct {
	Get_contract_request Get_contract_request
	Return_channel       chan echonetlite.Echonetlite
}

type Data_context struct {
	Post_data_request Post_data_request
}

//////////////////////////////////////////////

// http request body structure
type Get_contract_request struct {
	Gw_id string `json:"gw_id"`
}

type Post_data_request struct {
	Gw_id  string `json:"gw_id"`
	Format string `json:"format"`
}

type Get_contract_response struct {
	Format string `json:"format"`
}

/////////////////////////////////////

type Echonet_instance struct {
	Read_recv_contract chan Contract_context
	Read_recv_data     chan Data_context
}

func (self *Echonet_instance) Contract(w http.ResponseWriter, r *http.Request) {
	length := r.ContentLength
	reqBody := make([]byte, length)
	var ctx Get_contract_request
	err := json.Unmarshal(reqBody, &ctx)
	if err != nil {
		fmt.Println(err)
		fmt.Println("json.Unmarshal error")
		return
	}
	var recv_context Contract_context = Contract_context{ctx, make(chan echonetlite.Echonetlite)}
	self.Read_recv_contract <- recv_context
	echonetlite := <-recv_context.Return_channel

	format := Get_contract_response{Format: string(echonetlite.Frame[:echonetlite.Frame_size])}
	json_buf, err := json.Marshal(format)
	if err != nil {
		fmt.Println(err)
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

func (self *Echonet_instance) Data(w http.ResponseWriter, r *http.Request) {
	length := r.ContentLength
	reqBody := make([]byte, length)
	var ctx Post_data_request
	err := json.Unmarshal(reqBody, &ctx)
	if err != nil {
		fmt.Println(err)
		fmt.Println("json.Unmarshal error")
		return
	}
	var recv_context Data_context = Data_context{ctx}
	self.Read_recv_data <- recv_context

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func Init(addr, port, contract, data string) Echonet_instance {
	var echonet_instance Echonet_instance = Echonet_instance{make(chan Contract_context), make(chan Data_context)}
	server := http.Server{
		Addr: addr + ":" + port,
	}

	http.HandleFunc(contract, echonet_instance.Contract)
	http.HandleFunc(data, echonet_instance.Data)

	go server.ListenAndServe()

	return echonet_instance
}
