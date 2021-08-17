package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func sendJsonPost() {
	url := "http://localhost:8080/TodoPostJson" //请求地址
	contentType := "application/json"
	//参数，多个用&隔开
	data := strings.NewReader("username=admin&&password=123456")
	resp, err := http.Post(url, contentType, data)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}
