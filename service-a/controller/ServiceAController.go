package controller

import (
	"fmt"
	"github.com/lengrongfu/telepresence-demo/service-a/config"
	"io/ioutil"
	"net/http"
)

func Demo(w http.ResponseWriter, r *http.Request) {
	//log.LOG.Infof("this is request ip is:", r.RemoteAddr)
	w.Write([]byte("this is A service."))
	return
}

func CallServiceAll(w http.ResponseWriter,r * http.Request) {
	resp, err := http.Get(config.Conf.ServiceB + "/B/service-b")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.Write([]byte(fmt.Errorf("read service-b input bytes error %s",err.Error()).Error()))
		return
	}
	fmt.Printf("read service B %s interface result %s\n",r.RequestURI,string(bytes))
}

func CallServiceB(w http.ResponseWriter,r * http.Request)  {
	resp, err := http.Get(config.Conf.ServiceB + "/B/demo-b")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.Write([]byte(fmt.Errorf("read service-b input bytes error %s",err.Error()).Error()))
		return
	}
	fmt.Printf("read service B /B/demo-b interface result %s\n",string(bytes))
}


func CallServiceC(w http.ResponseWriter,r * http.Request)  {
	resp, err := http.Get(config.Conf.ServiceC + "/C/demo-c")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.Write([]byte(fmt.Errorf("read service-c input bytes error %s",err.Error()).Error()))
		return
	}
	fmt.Printf("read service c /C/demo-c interface result %s\n",string(bytes))
}