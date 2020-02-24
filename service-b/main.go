package main

import (
	"fmt"
	"github.com/lengrongfu/telepresence-demo/service-b/config"
	"github.com/lengrongfu/telepresence-demo/service-b/controller"
	"net/http"
	"strconv"
)

func main() {
	config.Parse()

	fmt.Println("service b start ....")
	http.HandleFunc("/B/demo", controller.Demo)
	http.HandleFunc("/B/demo-b",controller.DemoB)
	http.HandleFunc("/B/service-b",controller.CallServiceB)
	http.HandleFunc("/B/service-a",controller.CallServiceA)
	http.ListenAndServe(":"+strconv.Itoa(config.Conf.Port), nil)

}
