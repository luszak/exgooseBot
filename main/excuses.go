package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Excuse struct {
	Id     int    `json:"id"`
	Excuse string `json:"excuse"`
}

func decodeExcuse(r *http.Response) Excuse {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("Error while reading the body: " + err.Error())
	}
	var exc Excuse
	err = json.Unmarshal(body, &exc)
	if err != nil {
		panic("Error while converting the body: " + err.Error())
	}
	return exc
}

func requestExcuse() string {
	resp, err := http.Get("https://theexcusegoose.com/generate/")
	if err != nil {
		panic("Error while invoking a request: " + err.Error())
	}
	excuse := decodeExcuse(resp)
	return excuse.Excuse
}

func GetExcuse() string {
	return "I can't because " + requestExcuse()
}
