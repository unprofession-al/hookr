package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/apex/gateway"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	yaml "gopkg.in/yaml.v2"
)

// Sever holds all dependencies of the webserver
type Server struct {
	listener string
	handler  http.Handler
	hooks    Hooks
}

func NewServer(listener, configFile string) (Server, []error) {
	s := Server{
		listener: listener,
	}

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return s, []error{fmt.Errorf("could not read file '%s', error was %s", configFile, err)}
	}

	var hooks Hooks
	err = yaml.Unmarshal(yamlFile, &hooks)
	if err != nil {
		return s, []error{fmt.Errorf("could not unmarshal file '%s', error was %s", configFile, err)}
	}

	errs := hooks.Prepare()
	if len(errs) > 0 {
		return s, errs
	}
	s.hooks = hooks

	r := mux.NewRouter().StrictSlash(true)

	routes := s.routes()
	routes.Populate(r, "")

	s.handler = alice.New(s.LoggerMiddleware).Then(r)
	return s, []error{}
}

func (s Server) run() {
	if s.listener != "" {
		fmt.Printf("Serving at http://%s\nPress CTRL-c to stop...\n", s.listener)
		log.Fatal(http.ListenAndServe(s.listener, s.handler))
	} else {
		fmt.Printf("Serving as lambda...\n")
		log.Fatal(gateway.ListenAndServe(s.listener, s.handler))
	}
}

func (s Server) respond(res http.ResponseWriter, req *http.Request, code int, data interface{}) {
	var err error
	var errMesg []byte
	var out []byte

	f := req.Header.Get("Accept")
	if f == "text/yaml" {
		res.Header().Set("Content-Type", "text/yaml; charset=utf-8")
		out, err = yaml.Marshal(data)
		errMesg = []byte("--- error: failed while rendering data to yaml")
	} else {
		res.Header().Set("Content-Type", "application/json; charset=utf-8")
		out, err = json.MarshalIndent(data, "", "    ")
		errMesg = []byte("{ 'error': 'failed while rendering data to json' }")
	}

	if err != nil {
		out = errMesg
		code = http.StatusInternalServerError
	}
	res.WriteHeader(code)
	res.Write(out)
}

func (s Server) raw(res http.ResponseWriter, code int, data []byte) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(code)
}
