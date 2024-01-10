package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.sample/rds"
)

func main() {

	var contentBytes []byte
	var err error
	var f *os.File

	f, err = os.OpenFile("./logs/web.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	log.Println("Hello, you may see this in log file")

	rds.InitRedis("redis:6379", "", 0)

	cfgFilePath := "./cfg/cfg.json"

	if contentBytes, err = os.ReadFile(cfgFilePath); err != nil {
		panic(fmt.Sprintf("load config file: %s error %v", cfgFilePath, err.Error()))
	}

	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.LoadHTMLFiles("./cfg/hello.tmpl")

	router.GET("/", func(c *gin.Context) {
		log.Println("router / got hit")
		c.HTML(http.StatusOK, "hello.tmpl", gin.H{
			"show": string(contentBytes),
		})
	})

	log.Println("router listening on :8081")
	router.Run(":8081")

}
