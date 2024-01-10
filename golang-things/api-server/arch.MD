## 常見系統串接架構
* 保持連線、同步程序
* 不保持連線、同步程序
* 保持連線、非同步程序 - server and client\polling 
* 不保持連線、非同步程序 - api and webhook 

### Part 1 - 保持連線、同步程序
* http client <-> http server
* 需考慮風險： 連線逾時、斷線

### Part 2 - 不保持連線、同步程序
* http client -> http server
* http server -> data store
* http client <-> data store
* 需考慮風險： data store、使用較多程序與系統資源

### Part 3 - 保持連線、非同步程序
* socket client <-> socket server
* 連線、斷開、傳送、接收為非同步進行。
* 匿名存取、容與丟失訊息、multitasking
* 需考慮風險： server down、連線管理

### Part 4 - 不保持連線、非同步程序
* http server request -> http server serve
* http server webhook <- http server serve
* 需考慮風險： token


