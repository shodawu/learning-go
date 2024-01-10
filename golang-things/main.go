package main

import (
	"flag"
	"fmt"
	"number-game/consts"
)

var nSrvCli = flag.Int("n", 1, "1: run server, 2: run client.")

func main() {
	flag.Parse()
	fmt.Printf("hello, world n= %v\n", *nSrvCli)

	// lstSample := hubs.Controller{}
	// lstSample.Hubs = game.LoadSamples()
	// for _, cmd := range lstSample.Hubs {
	// 	ok := cmd.Exec()
	// 	log.Println("testing passed: ", ok)
	// }

	// svcnotify.TriggerExec()

	// switch *nSrvCli {
	// case 1:
	// 	fmt.Println("Run Server")
	// 	apiserver.RunServer()
	// case 2:
	// 	fmt.Println("Run Clients")
	// 	webclient.RunClientP4()
	// }

	fmt.Println(consts.Ver)
}
