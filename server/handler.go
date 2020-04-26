package server

import (
	"fmt"
	"github.com/alecthomas/log4go"
	"io/ioutil"
	"net/http"
)

var (
	log *log4go.Logger
)

// initialize server by host and port
func InitServer(host, port string, logger *log4go.Logger) {
	log = logger
	log.Info("<----- Starting server ----->")
	http.HandleFunc("/", handler)
	_ = http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
	log.Info("<----- Server stopped ----->")
}

// handle messages
func handler(w http.ResponseWriter, resp *http.Request) {
	log.Info("Received message")
	_, _ = fmt.Fprintf(w, "%s %s %s\n", resp.Method, resp.URL, resp.Proto)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := string(body)
	_, _ = fmt.Fprintln(w, result)
	log.Info("Received body:\n" + result)
}
