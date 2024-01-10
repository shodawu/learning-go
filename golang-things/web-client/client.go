package webclient

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// RunClient ...
func RunClient() {

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(iProcess int) {
			fmt.Println("ProcessID: ", iProcess)
			defer wg.Done()
			for {
				r, err := http.Post("http://localhost:1234/ping", "application/json", nil)

				if err != nil {
					fmt.Println(err)
					time.Sleep(5 * time.Second)
					continue
				}
				defer r.Body.Close()
				body, err := ioutil.ReadAll(r.Body)
				fmt.Println(string(body))
			}
		}(i)
	}
	wg.Wait()

}

func RunClientP2() {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(iProcess int) {
			fmt.Println("ProcessID: ", iProcess)
			defer wg.Done()
			cFileName := make(chan string)
			cUUID := make(chan string)

			go func(c chan string) {
				var err = errors.New("not conneted")
				var r *http.Response
				for err != nil {
					r, err = http.Post("http://localhost:1234/ping", "application/json", nil)
					if err == nil {
						break
					}
					fmt.Println(err)
					time.Sleep(5 * time.Second)
				}
				defer r.Body.Close()
				body, err := ioutil.ReadAll(r.Body)
				fileName := fmt.Sprintf("%v.log", string(body))
				fmt.Println(iProcess, "API returns", fileName)

				go p2UUIDChan(c, fileName)
			}(cFileName)

			go func() {
				s := <-cUUID
				fmt.Println(iProcess, "Find log", s)
			}()

			select {
			case fName := <-cFileName:
				var erFile = errors.New("not exists")
				var f *os.File
				for erFile != nil {
					f, erFile = os.Open(fName)
					if erFile != nil {
						time.Sleep(100 * time.Millisecond)
					}
				}
				b, _ := ioutil.ReadAll(f)
				go p2UUIDChan(cUUID, string(b))
			}

		}(i)
	}
	wg.Wait()

}

func p2UUIDChan(c chan string, data string) {
	c <- data
}

type Conn struct {
	ID         int
	SocketConn *websocket.Conn
	Recv       chan []byte
}

func RunClientP3() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		u := url.URL{Scheme: "ws", Host: "localhost:1234", Path: "ping"}
		reqHeader := http.Header{}
		wsConn, _, err := websocket.DefaultDialer.Dial(u.String(), reqHeader)

		if err == nil {
			fmt.Printf("Process: %v connected\n", i)
			cli := Conn{
				ID:         i,
				SocketConn: wsConn,
			}
			go cli.Listen()
		}
	}
	wg.Wait()
}

func (c *Conn) Listen() {
	defer func() {
		c.SocketConn.Close()
		fmt.Println(c.ID, "was disconnected")
	}()

	for {
		_, message, err := c.SocketConn.ReadMessage()
		if err != nil {
			fmt.Printf("Process: %v, websocket error: %v\n", c.ID, err)
			c.SocketConn.Close()
			break
		}
		fmt.Printf("Process: %v, Get Message: %v\n", c.ID, string(message))

	}
}

func RunClientP4() {
	s := &http.Server{
		Addr:           ":2234",
		Handler:        http.HandlerFunc(handlePong),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		for {
			r, err := http.Post("http://localhost:1234/ping", "application/json", nil)
			if err != nil {
				fmt.Println(err)
				time.Sleep(5 * time.Second)
				continue
			}
			fmt.Println("Process send request")
			defer r.Body.Close()
			time.Sleep(500 * time.Millisecond)
		}
	}()

	erPong := s.ListenAndServe()
	fmt.Println("Listen :2234", erPong)

}

func handlePong(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
}
