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
    format string `json:"format"`
}
```

Post request data/

```
type Post_data_request struct{
    gw_id string `json:"gw_id"`
    format string `json:"format"`
}
```
