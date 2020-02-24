package main

import (
	"fmt"
	"github.com/lengrongfu/telepresence-demo/service-c/config"
	"github.com/lengrongfu/telepresence-demo/service-c/controller"
	"net/http"
	"strconv"
)

func main() {
	config.Parse()

	fmt.Println("service c start ....")
	http.HandleFunc("/C/demo", controller.Demo)
	http.HandleFunc("/C/demo-c",controller.DemoC)
	http.HandleFunc("/C/service-c",controller.CallServiceC)
	http.ListenAndServe(":"+strconv.Itoa(config.Conf.Port), nil)

}
