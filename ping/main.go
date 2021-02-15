package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/lockwooddev/practical-k8s/utils"
	"github.com/sirupsen/logrus"
)

var hostname string
var port string

func init() {
	hostname = utils.GetEnv("HOSTNAME", true)
	port = utils.GetEnv("PORT", true)
}

var exit = make(chan bool)

func main() {
	logrus.Info("starting ping service")
	go pinger()
	<-exit
}

func pinger() error {
	for {
		time.Sleep(time.Second * 2)

		url := fmt.Sprintf("http://%s:%s/ping", hostname, port)
		logrus.Infof("request to %s", url)
		res, err := http.Get(url)
		if err != nil {
			logrus.Warnf("trouble reaching pong service: %s", err.Error())
			continue
		}

		if res.StatusCode != 200 {
			logrus.Warnf("pong service invalid status code %d", res.StatusCode)
			continue
		}

		content, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return err
		}
		logrus.Infof("response: %d - %s", res.StatusCode, string(content))
	}
}
