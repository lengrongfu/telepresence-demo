package controller

import (
	"fmt"
	"github.com/lengrongfu/telepresence-demo/service-a/config"
	"github.com/lengrongfu/telepresence-demo/service-c/log"
	"io/ioutil"
	"net/http"
)

func Demo(w http.ResponseWriter, r *http.Request) {
	//log.LOG.Infof("this is request ip is:", r.RemoteAddr)
	w.Write([]byte("this is C service."))
	return
}

func DemoC(w http.ResponseWriter, r *http.Request) {
	log.LOG.Infof("this is service-c request ip is:", r.RemoteAddr)
	w.Write([]byte("this is c service."))
	return
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
	fmt.Fprint(w,fmt.Sprintf("read service c %s interface result %s\n",r.RequestURI,string(bytes)))

}