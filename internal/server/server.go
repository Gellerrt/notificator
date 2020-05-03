package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Server struct {
	version string
	conf    *Config
	logger  *logrus.Logger
	router  *mux.Router
}

func New(c *Config) *Server {
	return &Server{
		version: "0.0.1",
		conf:    c,
		logger:  logrus.New(),
		router:  mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	if err := s.configLogger(); err != nil {
		return err
	}
	s.configRouter()
	s.logger.Info(fmt.Sprintf("<-------------------- Application version %s started -------------------->", s.version))
	return http.ListenAndServe(fmt.Sprintf("%s:%s", s.conf.Host, s.conf.Port), s.router)
}

func (s *Server) configLogger() error {
	level, err := logrus.ParseLevel(s.conf.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	s.logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     false,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})
	f, err := os.OpenFile(s.conf.LogName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	s.logger.SetOutput(f)
	return nil
}

func (s *Server) configRouter() {
	s.router.HandleFunc(s.conf.URI, s.handleNotificate())
}

func (s *Server) handleNotificate() http.HandlerFunc {
	return func(w http.ResponseWriter, resp *http.Request) {
		s.logger.Info("Received message")
		fmt.Fprintf(w, "%s %s %s\n", resp.Method, resp.URL, resp.Proto)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		result := string(body)
		fmt.Fprintln(w, result)
		s.logger.Info("Received body:\n" + result)
	}
}
