## echonetlite http bridge client package

Get request contract/

```
type Get_cnotract_request struct{
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

### Client Start up

```
Init(addr, port, contract_path, data_path string)
loop{
    Contract()
    go Data()
}
```
