package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	addr = ""
	port = "8080"
)

type Get_echonet_request struct {
	gw_id  string `json:"gw_id"`
	format string `json:"format"`
}

type Get_echonet_response struct {
	format string `json:"format"`
}

type Echonet_bridge_server struct{
	channel chan Get_echonet_request
}

type Bridge_interface struct{

}

var bridge_server Echonet_bridge_server

func Handler(w http.ResponseWriter, r *http.Request) {
	length := r.ContentLength
	reqBody := make([]byte, length)
	var ctx Get_echonet_request
	err := json.Unmarshal(reqBody, &ctx)
	if err != nil {
		fmt.Println("json.Unmarshal error")
		return
	}
	fmt.Println(ctx)

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

func Init(addr, port string) {
	bridge_server = make(chan )
	server := http.Server{
		Addr: addr + ":" + port,
	}

	http.HandleFunc("/echonet", Handler)

	go server.ListenAndServe()
}
