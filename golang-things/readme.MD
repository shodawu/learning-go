# golang 工具介紹
## go run 
* go run main.go
## go env
## go mod
* go mod init number-game 
## go test
* go test ./game
* go test ./... 
## go build
* GOOS=linux GOARCH=386 go build
* GOOS=windows GOARCH=amd64 go build -o ./number-guess
### build constraints
* // +build
* 於.go檔案中出現在`package`上方，且下方需要再完全空一行
* `,` 意義為 AND。 ` `意義為 OR。 `!`意義為 NOT
* `linux,386 darwin,!cgo` 意思等同 `(linux AND 386) OR (darwin AND (NOT cgo))`
* 多行之間為AND關聯。如下方範例
```
// +build linux darwin
// +build amd64

等於 (linux OR darwin) AND amd64
```
* https://golang.org/cmd/go/#hdr-Build_constraints

## 參考資料
* https://golang.org/doc/install/source

# BDD behavior-driven development
* Gherkin語法：https://cucumber.io/docs/gherkin/reference/
* cucumber支援各種主流語言 Ruby、JAVA、javascript、.NET、PHP
* cucumber in golang: https://github.com/cucumber/godog
* .feature檔案方便閱讀、傳遞知識
* [範例](./game/features/game.feature)

# C4-Model範例
* [需求](./svcnotify/docs/req.md)
* [C4-Model](./svcnotify/docs/c4.md)

# 常見系統串接架構
* [4種串接架構介紹](./api-server/arch.MD) 

