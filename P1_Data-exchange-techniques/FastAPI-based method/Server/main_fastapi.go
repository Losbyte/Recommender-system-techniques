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
	if resp.StatusCode != 200 {
		// FastAPI错误返回一般是JSON格式： {"detail": "错误描述"}
		var errResp map[string]interface{}
		err := json.Unmarshal(body, &errResp)
		if err == nil {
			fmt.Printf("API Error: %v\n", errResp["detail"])
		} else {
			fmt.Printf("API returned status %d, body: %s\n", resp.StatusCode, string(body))
		}
		return
	}

	var resData AddResponse
	err = json.Unmarshal(body, &resData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result from FastAPI: %d\n", resData.Result)
}
