package apiserver

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	webclient "number-game/web-client"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var mu sync.Mutex
var cUUID chan string
var mP3Conns map[*webclient.Conn]bool

// RunServer ...
func RunServer() {

	cUUID = make(chan string)
	// go storeUUID2(cUUID)
	go storeUUID4(cUUID)

	// mP3Conns = make(map[*webclient.Conn]bool)
	// go storeUUID3()

	s := &http.Server{
		Addr:           ":1234",
		Handler:        http.HandlerFunc(handlePing4),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
func handlePing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getUUID())
}

func getUUID() string {
	mu.Lock()
	time.Sleep(500 * time.Millisecond)
	u, _ := uuid.NewUUID()
	mu.Unlock()
	return u.String()

}

func handlePing2(w http.ResponseWriter, r *http.Request) {
	u, _ := uuid.NewUUID()
	go func(c chan string, data string) {
		c <- data
	}(cUUID, u.String())

	fmt.Fprintf(w, u.String())
}

func storeUUID2(c chan string) {
	for {
		s := <-c
		time.Sleep(2000 * time.Millisecond)
		ioutil.WriteFile(fmt.Sprintf("%v.log", s), []byte(s), 0775)
	}
}

func handlePing3(w http.ResponseWriter, r *http.Request) {
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		}}).Upgrade(w, r, nil)

	c := webclient.Conn{
		SocketConn: conn,
	}
	if err == nil {
		mP3Conns[&c] = true
	}

}

func storeUUID3() {
	for {
		for c := range mP3Conns {
			u, _ := uuid.NewUUID()
			c.SocketConn.WriteMessage(websocket.TextMessage, []byte(u.String()))
		}
		time.Sleep(5000 * time.Millisecond)
	}
}

func handlePing4(w http.ResponseWriter, r *http.Request) {
	u, _ := uuid.NewUUID()
	go func(c chan string, data string) {
		c <- data
	}(cUUID, u.String())
}

func storeUUID4(c chan string) {
	for {
		s := <-c
		time.Sleep(2000 * time.Millisecond)
		_, err := http.Post("http://localhost:2234/pong", "application/json",
			bytes.NewBuffer([]byte(s)))
		if err != nil {
			fmt.Println("response error", err)
		}
	}
}
