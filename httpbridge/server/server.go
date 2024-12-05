package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hasuburero/echonetlite/echonetlite"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var Sig chan os.Signal

// data structure witch through the api channel
type Contract_context struct {
	Get_contract_request Get_contract_request
	Return_channel       chan echonetlite.Echonetlite
}

type Data_context struct {
	Post_data_request Post_data_request
}

// http request body structure
type Get_contract_request struct {
	Gw_id string `json:"gw_id"`
}

type Post_data_request struct {
	Gw_id string `json:"gw_id"`
	Frame string `json:"frame"`
}

type Get_contract_response struct {
	Frame string `json:"frame"`
}

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

	frame := Get_contract_response{Frame: string(echonetlite.Frame[:echonetlite.Frame_size])}
	json_buf, err := json.Marshal(frame)
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
	Sig = make(chan os.Signal)
	signal.Notify(Sig, syscall.SIGTERM)
	var echonet_instance Echonet_instance = Echonet_instance{make(chan Contract_context), make(chan Data_context)}
	server := http.Server{
		Addr: addr + port,
	}

	http.HandleFunc(contract, echonet_instance.Contract)
	http.HandleFunc(data, echonet_instance.Data)

	go func() {
		err := server.ListenAndServe()
		if err != nil {

		}
	}()

	return echonet_instance
}
