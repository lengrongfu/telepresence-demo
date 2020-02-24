package main

import (
	"fmt"
	"github.com/lengrongfu/telepresence-demo/service-a/controller"
	"github.com/lengrongfu/telepresence-demo/service-a/config"
	"net/http"
	"strconv"
)





func main() {
	config.Parse()

	fmt.Println("service a start....")
	http.HandleFunc("/A/demo", controller.Demo)
	http.HandleFunc("/A/service-b",controller.CallServiceB)
	http.HandleFunc("/A/service-c",controller.CallServiceC)
	http.HandleFunc("/A/all",controller.CallServiceAll)
	http.ListenAndServe(":"+strconv.Itoa(config.Conf.Port), nil)

}
