package controller

import (
	"fmt"
	"github.com/lengrongfu/telepresence-demo/service-b/config"
	"github.com/lengrongfu/telepresence-demo/service-b/log"
	"io/ioutil"
	"net/http"
)

func Demo(w http.ResponseWriter, r *http.Request) {
	//log.LOG.Infof("this is request ip is:", r.RemoteAddr)
	w.Write([]byte("this is B service."))
	return
}

func DemoB(w http.ResponseWriter, r *http.Request) {
	log.LOG.Infof("this is service-b request ip is:", r.RemoteAddr)
	w.Write([]byte("this is B service."))
	return
}


func CallServiceB(w http.ResponseWriter,r * http.Request)  {
	resp, err := http.Get(config.Conf.ServiceC + "/C/service-c")
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
	fmt.Printf("read service b %s interface result %s\n",r.RequestURI,string(bytes))
	fmt.Fprint(w,fmt.Sprintf("read service b %s interface result %s\n",r.RequestURI,string(bytes)))
}



func CallServiceA(w http.ResponseWriter,r * http.Request)  {
	resp, err := http.Get(config.Conf.ServiceA + "/A/demo")
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
	fmt.Printf("call service A result %s\n",string(bytes))
	fmt.Fprint(w,fmt.Sprintf("call service A result %s\n",string(bytes)))
}