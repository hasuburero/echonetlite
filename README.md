# echonetlite

## 参考資料

<https://echonet.jp/wp/wp-content/uploads/pdf/General/Standard/ECHONET_lite_V1_14_jp/ECHONET-Lite_Ver.1.14(02).pdf>

### tcpbridge

sample program  
NATを介して通信するためにclientからserverに対してtcpコネクションを作成  
コネクションを通してechonet lite frameを伝送

### httpbridge

sample program  
NATを介して通信するためにclientからmec-rmに対してhttp通信  
clientからの最初の通信はmec-rmに対してリクエスト(便宜上gw_idのみ)を送信  
mec-rmからのレスポンスでデータ要求・制御指示を送信  
clientからのリクエストでデータを返送・レスポンス(制御指示)待ち
