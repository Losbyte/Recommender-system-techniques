package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AddRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type AddResponse struct {
	Result int    `json:"result"`
	Error  string `json:"error,omitempty"`
}

func main() {
	url := "http://localhost:5000/add"

	reqData := AddRequest{A: 15, B: 25}
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var resData AddResponse
	err = json.Unmarshal(body, &resData)
	if err != nil {
		panic(err)
	}

	if resData.Error != "" {
		fmt.Printf("API Error: %s\n", resData.Error)
		return
	}

	fmt.Printf("Result from Flask API: %d\n", resData.Result)
}
