package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	//"github.com/hasuburero/echonetlite/echonetlite"
	"io"
	"net/http"
)

const (
	Get_method  = "GET"
	Post_method = "POST"
)

type Get_contract_request struct {
	Gw_id string `json:"gw_id"`
}

type Get_contract_response struct {
	Frame []byte `json:"frame"`
}

type Post_data_request struct {
	Gw_id string `json:"gw_id"`
	Frame []byte `json:"frame"`
}

type GW_instance struct {
	Gw_id         string
	Scheme        string
	Addr          string
	Port          string
	Contract_path string
	Data_path     string
}

var wait chan bool

func (self *GW_instance) Data(frame []byte) error {
	str := base64.StdEncoding.EncodeToString(frame)
	request := Post_data_request{Gw_id: self.Gw_id, Frame: []byte(str)}
	json_buf, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	buf := bytes.NewBuffer(json_buf)
	req, err := http.NewRequest(Post_method, self.Scheme+self.Addr+self.Port+self.Data_path, buf)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req.Header.Add("content-type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (self *GW_instance) Contract() ([]byte, error) {
	request := Get_contract_request{Gw_id: self.Gw_id}
	json_buf, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// debug
	buf := bytes.NewBuffer(json_buf)
	req, err := http.NewRequest(Get_method, self.Scheme+self.Addr+self.Port+self.Contract_path, buf)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("content-type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var ctx Get_contract_response
	err = json.Unmarshal(body, &ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	byte_buf, err := base64.StdEncoding.DecodeString(string(ctx.Frame))

	return byte_buf, nil
}

func Init(Gw_id, Scheme, Addr, Port, Contract_path, Data_path string) GW_instance {
	gw := GW_instance{Gw_id: Gw_id, Scheme: Scheme, Addr: Addr, Port: Port,
		Contract_path: Contract_path,
		Data_path:     Data_path}
	return gw
}
