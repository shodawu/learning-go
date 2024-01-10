# 後端工程師工作簡介

## 後端程式介紹
後端程式通常會負責儲存、轉換資料。<br>
讀寫持久性的資料，例如資料庫、硬碟檔案。<br>
讀寫暫時性的資料，但是要加快速度，就是快取。<br>
跟其他程序溝通，接收命令，或是發出請求。<br>
通常後端程式不會需要所謂的人機介面，就是渲染出一個很漂亮的畫面，展示資料。<br>
而是著重在於程式與程式之間資料傳遞與保存，為了效率，資料格式通常是二進制的`byte`型態。<br>
但為了方便除錯、監測，也會採用字串、JSON等方便人類理解的資料型態。<br>
以下介紹通常用後端程式使用哪些工具，處理哪些問題。<br>


## 後端程式處理的問題
* 程式啟動時，給予參數
* 程式啟動或執行時，讀取參數資料
* 程式執行時，透過TCP/IP協定傳輸
* 程式執行時，將資料儲存
* 程式執行時，傳出資料

## go處理以上問題時，常用的工具、套件
* 程式啟動時，給予參數
    * flag
    * os.Getenv
* 程式啟動或執行時，讀取參數資料
    * io.Reader
    * goini: `https://github.com/zieckey/goini`
* 程式執行時，透過TCP/IP、FTP等協定傳輸：
    * gin: `https://github.com/gin-gonic/gin`
    * ftp: `github.com/jlaffaye/ftp`
* 程式執行時，將資料儲存
    * gorm: `https://github.com/go-gorm/gorm`
    * reids: `https://github.com/redis/go-redis`
    * mongodb: `https://github.com/mongodb/mongo-go-driver`
    * google drive: `google.golang.org/api/drive/v3`
    * google sheet: `google.golang.org/api/sheets/v4`
* 程式執行時，將資料轉檔
    * zip
    * excel: `https://github.com/qax-os/excelize`


## 以使用flag套件為例，練習撰寫第一個go程式
1. flag套件使用說明：`https://pkg.go.dev/flag`
2. 建議一個程式資料夾，並使用go的套件管理`go mod`，以及使用git做版控
```
mkdir less-1 && cd less-1
go mod init less1
git init
```
3. 複製下面sample code，產生一個main.go檔案
```go
package main

import (
	"flag"
	"log"
)

var species = flag.String("species", "gopher", "the species we are studying")

func main() {
	flag.Parse()

	log.Printf("Hello World!! %v", *species)
}
```
 
4. 執行go程式，叫出flag的使用說明
```
go run main.go -h
```

5. 執行go程式，不傳入參數
```
go run main.go
```

6. 執行go程式，傳入任意參數
```
go run main.go -species hello
```