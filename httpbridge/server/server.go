package server

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/hasuburero/echonetlite/echonetlite"
	"io"
	"net/http"
	"time"
)

var Error = make(chan error, 1)

// data structure witch through the api channel
type Bridge_instance struct {
	Read_recv_contract chan Contract_context
	Read_recv_data     chan Data_context
}

type Contract_context struct {
	Gw_id          string
	Timestamp      time.Time
	Return_channel chan ReturnChannel
}

type Data_context struct {
	Gw_id     string
	Frame     string
	Timestamp time.Time
}

type ReturnChannel struct {
	Echonet_instance chan echonetlite.Echonetlite
	StatusCode       int
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

func (self *Bridge_instance) Contract(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	length := r.ContentLength
	reqBody := make([]byte, length)
	r.Body.Read(reqBody)
	var ctx Get_contract_request
	err := json.Unmarshal(reqBody, &ctx)
	if err != nil {
		fmt.Println(err)
		fmt.Println("json.Unmarshal error")
		return
	}

	var return_channel = make(chan ReturnChannel)
	defer close(return_channel)
	var recv_context Contract_context = Contract_context{Gw_id: ctx.Gw_id, Timestamp: now, Return_channel: return_channel}
	self.Read_recv_contract <- recv_context
	return_ctx := <-recv_context.Return_channel

	if return_ctx.StatusCode == http.StatusOK {
		echonet_instance := <-return_ctx.Echonet_instance
		fmt.Println(echonet_instance)
		err = echonet_instance.MakeFrame()
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		fmt.Println(echonet_instance.Frame)
		frame := Get_contract_response{Frame: base64.StdEncoding.EncodeToString(echonet_instance.Frame[:echonet_instance.Frame_size])}

		json_buf, err := json.Marshal(frame)
		if err != nil {
			fmt.Println(err)
			fmt.Println("json.Marshal error")
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		body := bytes.NewBuffer(json_buf)
		_, err = io.Copy(w, body)
		if err != nil {
			fmt.Println("io.Copy error")
			w.WriteHeader(http.StatusBadGateway)
			return
		}
	} else {
		w.WriteHeader(return_ctx.StatusCode)
	}
}

func (self *Bridge_instance) Data(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	length := r.ContentLength
	reqBody := make([]byte, length)
	r.Body.Read(reqBody)
	var ctx Post_data_request
	err := json.Unmarshal(reqBody, &ctx)
	if err != nil {
		fmt.Println(err)
		fmt.Println("json.Unmarshal error")
		return
	}
	byte_buf, err := base64.RawStdEncoding.DecodeString(ctx.Frame)
	ctx.Frame = string(byte_buf)
	var recv_context Data_context = Data_context{Gw_id: ctx.Gw_id, Frame: ctx.Frame, Timestamp: now}
	self.Read_recv_data <- recv_context

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func Start(addr, port, contract, data string) Bridge_instance {
	var bridge_instance Bridge_instance = Bridge_instance{make(chan Contract_context), make(chan Data_context)}
	server := http.Server{
		Addr: addr + ":" + port,
	}

	http.HandleFunc(contract, bridge_instance.Contract)
	http.HandleFunc(data, bridge_instance.Data)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println(err)
			fmt.Println("http.Server.ListenAndServe error")
			Error <- err
		}
	}()

	return bridge_instance
}
