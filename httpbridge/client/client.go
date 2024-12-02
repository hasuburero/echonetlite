package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	Get_method  = "GET"
	Post_method = "POST"
)

type Get_contract_request struct {
	gw_id string `json:"gw_id"`
}

type Get_contract_response struct {
	frame string `json:"frame"`
}

type Post_data_request struct {
	gw_id string `json:"gw_id"`
	frame string `json:"frame"`
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

func (self *GW_instance) Data(frame string) error {
	request := Post_data_request{gw_id: self.Gw_id, frame: frame}
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))

	return nil
}

func (self *GW_instance) Contract() (string, error) {
	request := Get_contract_request{self.Gw_id}
	json_buf, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	buf := bytes.NewBuffer(json_buf)
	req, err := http.NewRequest(Get_method, self.Scheme+self.Addr+self.Port+self.Contract_path, buf)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	req.Header.Add("content-type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var ctx Get_contract_response
	err = json.Unmarshal(body, &ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return ctx.frame, nil
}

func Init(Gw_id, Addr, Port, Contract_path, Data_path string) GW_instance {
	gw := GW_instance{Gw_id: Gw_id, Addr: Addr, Port: Port,
		Contract_path: Contract_path,
		Data_path:     Data_path}
	return gw
}
