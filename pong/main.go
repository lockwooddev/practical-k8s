package main

import (
	"fmt"
	"net/http"

	"github.com/lockwooddev/practical-k8s/utils"
	"github.com/sirupsen/logrus"
)

func ping(w http.ResponseWriter, req *http.Request) {
	logrus.Infof("receiver /ping request from %s", req.RemoteAddr)
	fmt.Fprintf(w, "pong")
}

func main() {
	port := utils.GetEnv("PORT", true)

	logrus.Infof("starting pong service on port %s and waiting for /ping requests", port)
	http.HandleFunc("/ping", ping)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
