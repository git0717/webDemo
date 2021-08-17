package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"unsafe"
)

func main() {
	apiUrl := "http://127.0.0.1:8080" //请求地址
	//SamplePostJson(url)
	httpPostForm(apiUrl)
}
func httpGet(url string) {
	var r http.Request
	r.ParseForm()
	r.Form.Add("username", "zhangsan")
	r.Form.Add("password", "123456")
	bodystr := strings.TrimSpace(r.Form.Encode())
	request, err := http.NewRequest("GET", url, strings.NewReader(bodystr))
	if err != nil {
		//TODO:
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Connection", "Keep-Alive")

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)

}
func httpPostForm(apiUrl string) {
	resource := "/TodoPostForm"
	data := url.Values{}
	data.Set("userName", "zhangsan")
	data.Set("password", "123456")
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String() // "http://127.0.0.1/tpost"

	request, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		//TODO:
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Connection", "Keep-Alive")

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)

}
func SamplePostJson(url string) {
	song := make(map[string]interface{})
	song["username"] = "admin"
	song["timelength"] = 128
	song["password"] = "123456"
	bytesData, err := json.Marshal(song)
	if err != nil {
		fmt.Println(err.Error() )
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
}