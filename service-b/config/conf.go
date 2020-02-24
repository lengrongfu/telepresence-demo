package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Port int `json:"port"`
	ServiceB string `json:"service_b"`
	ServiceC string `json:"service_c"`
	ServiceA string `json:"service_a"`
}



var Conf Config

var filePath *string

func init() {
	filePath = flag.String("conf", "/config.json", "config file path")
	flag.Parse()

}

func Parse()  {
	fmt.Println("config file path is:", *filePath)

	data, err := ioutil.ReadFile(*filePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &Conf)
	if err != nil {
		panic(err)
	}
}