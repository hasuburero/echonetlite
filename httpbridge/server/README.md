## echonet lite http bridge pakcage

Get request contract/

```
type Get_contract_request struct{
    gw_id string `json:"gw_id"`
}
```

Get response contract/

```
type Get_contract_response struct{
    frame string `json:"frame"`
}
```

Post request data/

```
type Post_data_request struct{
    gw_id string `json:"gw_id"`
    frame string `json:"frame"`
}
```

### Bridge Start up

```
Init(addr, port, contract_path, data_path string) Echonet_instance
loop{
    <- Echonet_instance.read_recv_contract
    <- Echonet_instance.read_recv_data
}
```

```
type Echonet_instance struct{
    read_recv_contract chan Contract_context
    read_recv_data chan Data_context
}
type Contract_context struct{
    get_contract_request Get_contract_request
    return_channel chan echonetlite.Echonetlite
}
type Data_context struct{
    post_data_request Post_data_request
}
```
