package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func (s Server) HookHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("hook handler triggererd")
	vars := mux.Vars(req)
	hName, ok := vars["hook"]
	if !ok {
		s.respond(res, req, http.StatusNotFound, fmt.Sprintf("hook not provided"))
		fmt.Println(fmt.Sprintf("hook not provided, %v", vars))
		return
	}

	reqText, _ := json.Marshal(req)
	fmt.Println(reqText)

	body, err := ioutil.ReadAll(req.Body)
	fmt.Println(body)
	defer req.Body.Close()
	if err != nil {
		s.respond(res, req, http.StatusInternalServerError, fmt.Sprintf("could not read request body: %s", err.Error()))
		fmt.Println(fmt.Sprintf("Body was trash"))
		return
	}

	h, ok := s.hooks[hName]
	if !ok {
		s.respond(res, req, http.StatusNotFound, fmt.Sprintf("hook '%s' not defined", hName))
		fmt.Println(fmt.Sprintf("hook '%s' not defined", hName))
		return
	}

	err = h.Process(body)
	if err != nil {
		s.respond(res, req, http.StatusInternalServerError, fmt.Sprintf("could not process hook: %s", err.Error()))
		fmt.Println(fmt.Sprintf("could not process hook: %s", err.Error()))
		return
	}

	s.respond(res, req, http.StatusOK, "sent")
}

func (s Server) PingdomHookHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("hook handler triggererd")

	reqText, _ := json.Marshal(req)
	fmt.Println(reqText)

	body, err := ioutil.ReadAll(req.Body)
	fmt.Println(body)
	defer req.Body.Close()
	if err != nil {
		s.respond(res, req, http.StatusInternalServerError, fmt.Sprintf("could not read request body: %s", err.Error()))
		fmt.Println(fmt.Sprintf("Body was trash"))
		return
	}

	hName := "pingdomSwisstopoToSms"
	h, ok := s.hooks[hName]
	if !ok {
		s.respond(res, req, http.StatusNotFound, fmt.Sprintf("hook '%s' not defined", hName))
		fmt.Println(fmt.Sprintf("hook '%s' not defined", hName))
		return
	}

	err = h.Process(body)
	if err != nil {
		s.respond(res, req, http.StatusInternalServerError, fmt.Sprintf("could not process hook: %s", err.Error()))
		fmt.Println(fmt.Sprintf("could not process hook: %s", err.Error()))
		return
	}

	s.respond(res, req, http.StatusOK, "sent")
}
