# grpc-server
* https://grpc.io/docs/languages/go/quickstart/
* https://github.com/grpc/grpc-go

## install protobuf
```
brew install protobuf
protoc --version
```
## download golang packages
```
go get google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc \
    go get -u github.com/golang/protobuf/proto \
    google.golang.org/grpc
```
* 2023-03-12更新，如安裝時遇到`'go get' is no longer supported outside a module`錯誤訊息，需改以go install方式安裝。


* 如遇到錯誤訊息`protoc-gen-go: program not found or is not executable`，需以下列指令安裝`protoc-gen-go`
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

* 如遇到錯誤訊息`protoc-gen-go-grpc: program not found or is not executable`，需以下列指令安裝`protoc-gen-go-grpc`
```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## for schema first... generate gRPC code
### create .proto file
* https://github.com/grpc/grpc-go/blob/master/examples/helloworld/helloworld/helloworld.proto
* 複製以上範例檔，放在 /schema/hello.proto 路徑下。

### generate gRPC code command
```sh 
protoc --go_out=. --go-grpc_out=. schema/hello.proto
```
* 此命令會讓protoc讀取schema/hello.proto，依照`go_package`產生package路徑。 
* 並產生`_grpc.pb.go`、`.pb.go`兩個檔案，檔前綴名與`.proto`檔名相同。
* option go_package = 與 package 如果不一樣，rpc client啟動時會有錯誤訊息`rpc error: code = Unimplemented desc = method SayHello not implemented`。兩者設定為相同變正常了。

### gRPC service
* 於.proto檔了解此service定義了`SayHello` rpc Method，接收參數`HelloRequest`，返回`HelloReply`
* 打開 `greeter_server/main.go`實作`SayHello`。 
```go
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
        return &pb.HelloReply{Message: "Say Hello " + in.GetName()}, nil
}
```
* 打開 `greeter_client/main.go`實作呼叫。 
```go
r, err = c.SayHello(ctx, &pb.HelloRequest{Name: name})
if err != nil {
        log.Fatalf("could not greet: %v", err)
}
log.Printf("Greeting: %s", r.GetMessage())
```

### run server and client
* run server 
```
go run ./greeter_server/main.go
```
* run client
```
go run ./greeter_client/main.go Max
```

### work with database
* 將db connection instance在 greeter_server 的 `server`
* TODO: 或可嘗試instance在 `UnimplementedGreeterServer` 
